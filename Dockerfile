# Build stage
FROM golang:1.23-alpine AS builder
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
RUN apk update && apk add --no-cache git \
  build-base

WORKDIR /builder
COPY . .
RUN go build -o /builder/bin/goose /builder/cmd/databases/main.go && \
  go build -o /builder/bin/api /builder/cmd/app/main.go 

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /builder/bin/goose .
COPY --from=builder /builder/bin/api .
COPY --from=builder /builder/migrations ./migrations

CMD [ "./api" ]