FROM golang:1.19 AS dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 5000
CMD reflex go run cmd/server/main.go --start-service

FROM golang:1.19 AS build
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app cmd/server/main.go

FROM alpine:latest AS prod
RUN apk add --no-cache ca-certificates
COPY --from=build app .
EXPOSE 5000
CMD ./app