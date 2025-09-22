FROM node:22-alpine AS ui-builder

WORKDIR /app/ui

COPY ui/package*.json ./

RUN npm ci 

COPY ui/ ./

RUN npm run build

FROM golang:1.24-alpine AS go-builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY --from=ui-builder /app/ui/build ./ui/build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o r-place cmd/place/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /root/

COPY --from=go-builder /app/r-place .

RUN chown appuser:appgroup r-place

USER appuser

EXPOSE 9090 

CMD ["./r-place"]
