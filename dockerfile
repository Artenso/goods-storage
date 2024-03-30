FROM golang:1.22-alpine AS builder

COPY . /github.com/Artenso/goods-storage/
WORKDIR /github.com/Artenso/goods-storage/

RUN go mod download
RUN go build -o ./bin/goods-storage cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/Artenso/goods-storage/bin/goods-storage .

EXPOSE 50051
EXPOSE 8000

CMD ["./goods-storage"]