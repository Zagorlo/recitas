package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"
	"time"
)

func MakeContext() (context.Context, func()) {
	ctx, cancelCtx := context.WithCancel(context.TODO())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	errCh := make(chan error, 1)

	go func() {
		select {
		case <-sigCh:
			log.Printf("MakeContext cancelling\n")
			cancelCtx()
		case err := <-errCh:
			log.Printf("MakeContext fatal: %v\n", err)
			cancelCtx()
		}
	}()

	return ctx, cancelCtx
}

func MakeServer(ctx context.Context, cancelCtx func(), gracefulShutdownAwait time.Duration, handler http.Handler, addr string) {
	errCh := make(chan error, 1)

	var listener net.Listener
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	server := http.Server{Handler: handler}

	go func() {
		log.Printf("MakeServer start: %s\n", addr)
		errCh <- server.Serve(listener)
	}()

	select {
	case err = <-errCh:
		log.Printf("MakeServer fatal: %v\n", err)
		cancelCtx()
	case <-ctx.Done():
		log.Printf("MakeServer ctx done\n")
	}

	log.Printf("MakeServer closing\n")
	contextWithTimeout, _ := context.WithTimeout(context.TODO(), gracefulShutdownAwait)
	if err := server.Shutdown(contextWithTimeout); err != nil {
		log.Printf("MakeServer close error: %v\n", err)
	}
	log.Printf("MakeServer closed\n")
}

func ServeOrigin(allowed []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			stack := debug.Stack()
			_, _ = os.Stderr.Write(stack)
		}()

		origin := r.Header.Get("Origin")
		if len(origin) > 0 {
			if !isOriginAllowed(origin, allowed) {
				log.Printf("origin %v is not allowed", origin)
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte("origin not allowed"))
				return
			}
			header := w.Header()
			header.Add(acao, origin)
			header.Add(acam, allowMethods)
			header.Add(acah, allowHeaders)
			header.Add(acac, allowCredentials)
			header.Add(acma, maxAge)
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isOriginAllowed(origin string, allowed []string) bool {
	if len(allowed) == 0 {
		return true
	}
	u, err := url.Parse(origin)
	if err != nil {
		return false
	}
	h, _, _ := net.SplitHostPort(u.Host)
	if h == "" {
		h = u.Host
	}
	h = strings.ToLower(h)
	for _, it := range allowed {
		if h == it || strings.HasSuffix(h, "."+it) || it == "*" {
			return true
		}
	}
	return false
}

const (
	acao = "Access-Control-Allow-Origin"
	acam = "Access-Control-Allow-Methods"
	acah = "Access-Control-Allow-Headers"
	acac = "Access-Control-Allow-Credentials"
	acma = "Access-Control-Max-Age"

	allowHeaders     = "Origin, Authorization, Accept, Accept-Encoding, Cache-Control, Content-Type, Content-Length"
	allowMethods     = "GET, POST, PUT, DELETE, OPTIONS"
	allowCredentials = "true"
	maxAge           = "3600"
)
