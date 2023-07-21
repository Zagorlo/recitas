FROM golang:1.8.3-alpine3.6

RUN apk update && \
    apk add \
        bash \
        build-base \
        curl \
        make \
        git \
        && rm -rf /var/cache/apk/*

COPY . /go/src/recitas
WORKDIR /go/src/recitas

#RUN env GOOS="linux" go build -o ./cmd/main .
#RUN chmod 777 ./cmd/main

CMD ["./cmd/main"]
EXPOSE 8000