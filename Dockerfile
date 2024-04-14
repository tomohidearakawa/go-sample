FROM golang:1.22.1

COPY . /go/src/app
WORKDIR /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...

# アプリケーションを実行します。
CMD ["go-sample"]
