FROM golang:1.16 as build

COPY . /app
WORKDIR /app
RUN go build .

FROM alpine

COPY --from=build /app/db /app
ENTRYPOINT [ "/app/db" ]