FROM golang:1.22.1-alpine
RUN apk update && apk add git
WORKDIR /go/src
COPY . .

CMD ["go", "run", "main.go"]
