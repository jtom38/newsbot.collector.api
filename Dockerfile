FROM golang:1.18.2 as build

COPY . /app
WORKDIR /app
RUN go build .

FROM alpine

COPY --from=build /app/collector /app
ENTRYPOINT [ "/app/collector" ]