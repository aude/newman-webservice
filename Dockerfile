FROM golang:1.12-alpine AS builder
WORKDIR /app/
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o newman-webservice

FROM postman/newman:4
ENTRYPOINT ["/newman-webservice"]
COPY --from=builder /app/newman-webservice /newman-webservice
