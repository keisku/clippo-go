FROM golang:1.12.0-alpine3.9

WORKDIR /go/src/server_post

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh

EXPOSE 50052

CMD ["fresh"]

# FROM golang:1.12

# ENV PKG_PATH /go/src/github.com/kskumgk63/clippo-go
# ENV SERVICE_PATH $PKG_PATH/server_post
# ENV GO111MODULE on
# WORKDIR $PKG_PATH
# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# COPY . $PKG_PATH
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $SERVICE_PATH

# FROM alpine:3.8
# ENV TZ Asia/Tokyo
# COPY --from=0 /go/bin/server_post /go/bin/server_post
# ENTRYPOINT ["/go/bin/server_post"]
