FROM golang:1.18.4 as build

COPY . /app
WORKDIR /app
RUN go build .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest as app

RUN apk --no-cache add bash libc6-compat
RUN mkdir /app && mkdir /app/migrations
COPY --from=build /app/collector /app
COPY --from=build /go/bin/goose /app
COPY ./database/migrations/ /app/migrations

CMD [ "/app/collector" ]