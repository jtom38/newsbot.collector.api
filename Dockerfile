FROM golang:1.16 as build

COPY . /app
WORKDIR /app
RUN go build .

FROM alpine

COPY --from=build /app/collector /app
ENTRYPOINT [ "/app/collector" ]