FROM golang:1.21.1 AS buildstage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /main .


# Deploy stage

FROM alpine:latest

WORKDIR /

RUN apk add libc6-compat

COPY --from=buildstage /main /app/main

COPY .env /app/.env

EXPOSE 8000

ENTRYPOINT ["/app/main"]
