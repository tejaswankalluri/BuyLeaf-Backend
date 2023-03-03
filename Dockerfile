FROM golang:1.18-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build

# Create final image
FROM alpine
WORKDIR /
COPY --from=builder /build/myapp .
EXPOSE 8000
CMD ["./myapp"]