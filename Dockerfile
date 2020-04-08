FROM golang:1.14 as builder

WORKDIR /go/src/github.com/Alma-media/taxi/

ADD go.mod go.sum ./
RUN go mod download

ADD . /go/src/github.com/Alma-media/taxi/
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/taxi /go/src/github.com/Alma-media/taxi/main.go

FROM alpine:3.11.5

WORKDIR /app

COPY --from=builder /go/src/github.com/Alma-media/taxi/bin/taxi /app/

ENTRYPOINT ["/app/taxi"]
