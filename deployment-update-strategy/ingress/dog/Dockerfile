FROM golang:latest as builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v
RUN go build -v

FROM debian:stretch-slim
COPY --from=builder /go/src/app/app /app/
WORKDIR /app
EXPOSE 8080
CMD ["./app"]