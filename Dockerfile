FROM golang AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
ENV GO_ENV=production
ENTRYPOINT ["./app"]