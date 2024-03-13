FROM --platform=${BUILDPLATFORM} golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/my-tv-server ./cmd/server


FROM --platform=${TARGETPLATFORM} alpine:latest

WORKDIR /app

COPY --from=builder  /app/releases/my-tv-server .

RUN chmod +x ./my-tv-server

ENTRYPOINT ["./my-tv-server"]
CMD ["--port", "8000"]