# Build stage
FROM golang:1.26.0-alpine AS builder
RUN go install github.com/go-task/task/v3/cmd/task@latest
WORKDIR /chijji-moni
COPY app/go.mod app/go.sum ./
RUN go mod download
COPY . /chijji-moni/
RUN task build


# Run stage
FROM alpine:latest
WORKDIR /chijji-moni
COPY --from=builder /chijji-moni/dist/app .
EXPOSE 8080
CMD ["./app"]