FROM golang:1.20.2 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main-app ./

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main-app ./main-app
EXPOSE ${APP_PORT}

CMD ["./main-app"]