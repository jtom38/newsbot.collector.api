FROM golang:1.18.4 as build

COPY . /app
WORKDIR /app

# Always make sure that swagger docs are updated
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN /go/bin/swag i

# Always build the latest sql queries
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN /go/bin/sqlc generate

RUN go build .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest as app

RUN apk --no-cache add bash 
RUN apk --no-cache add libc6-compat
RUN apk --no-cache add chromium

RUN mkdir /app && mkdir /app/migrations
COPY --from=build /app/collector /app
COPY --from=build /go/bin/goose /app
COPY ./database/migrations/ /app/migrations

CMD [ "/app/collector" ]