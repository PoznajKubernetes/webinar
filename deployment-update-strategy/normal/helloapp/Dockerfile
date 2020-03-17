FROM golang:1.13-alpine AS build

WORKDIR /src

COPY main*.go go.* ./

RUN CGO_ENABLED=0 go build -o /bin/helloapp

FROM scratch

COPY --from=build /bin/helloapp /bin/helloapp

EXPOSE 8080

ENV environment=default

CMD ["/bin/helloapp"]