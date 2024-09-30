FROM golang:latest AS builder
#RUN apt-get update && apt-get install -y bash
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prometheus-demo .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/prometheus-demo .
ENTRYPOINT ["./prometheus-demo"]
CMD ["start"]