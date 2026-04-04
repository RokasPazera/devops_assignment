FROM golang:1.23-alpine AS builder
RUN apk add --no-cache gcc musl-dev
ENV CGO_ENABLED=1 GOTOOLCHAIN=auto
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /out/backend .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /out/backend /app/backend
CMD ["/app/backend"]
