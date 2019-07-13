FROM golang:1.12

WORKDIR /go/src/github.com/kskumgk63/clippo-go/server_cache
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]

# FROM golang:1.12

# ENV PKG_PATH /go/src/github.com/kskumgk63/clippo-go
# ENV SERVICE_PATH $PKG_PATH/server_cache
# ENV GO111MODULE on
# WORKDIR $PKG_PATH
# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# COPY . $PKG_PATH
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $SERVICE_PATH

# FROM alpine:3.8
# ENV TZ Asia/Tokyo
# COPY --from=0 /go/bin/server_cache /go/bin/server_cache
# ENTRYPOINT ["/go/bin/server_cache"]
