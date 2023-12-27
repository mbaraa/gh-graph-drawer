FROM golang:1.20-alpine as build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/github-graph-drawer ./run
COPY --from=build /app/templates ./templates

EXPOSE 8080

CMD ["./run"]
