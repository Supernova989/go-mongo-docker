FROM golang:1.16.0-alpine3.13
COPY . /app
WORKDIR /app
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux go build -o main
CMD ["./main"]
