FROM golang:1.12

ENV PKG_PATH /go/src/github.com/kskumgk63/clippo-go
ENV SERVICE_PATH $PKG_PATH
ENV GO111MODULE on
WORKDIR $PKG_PATH
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . $PKG_PATH
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $SERVICE_PATH

FROM alpine:3.8
ENV TZ Asia/Tokyo
COPY --from=0 /go/bin/front /go/bin/front
COPY ./static /static
COPY ./template /template
ENTRYPOINT ["/go/bin/front"]
