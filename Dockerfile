FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/jcf-go ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/jcf-go .

EXPOSE 8080

CMD ["./jcf-go"]
