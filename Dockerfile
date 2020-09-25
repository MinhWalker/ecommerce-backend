FROM alpine:3.12

RUN apk add --no-cache \
		ca-certificates && apk add golang && apk update

ENV GO111MODULE=on

ENV GOLANG_VERSION 1.15.2

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/ecommerce-backend

COPY . .

RUN go mod init ecommerce-backend

WORKDIR cmd/pro
RUN GOOS=linux go build -o app

ENTRYPOINT ["./app"]

EXPOSE 3000