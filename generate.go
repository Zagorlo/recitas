package recitas

//go:generate ./bin/protoc -I ./proto/include/ -I $GOPATH/src --proto_path=$GOPATH/src:. --validate_out=lang=go:$GOPATH/src --twirp_out=$GOPATH/src --go_out=$GOPATH/src ./proto/rec/rec.proto
