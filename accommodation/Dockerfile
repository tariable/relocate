FROM golang:1.16.5-alpine3.14 AS builder
RUN adduser -D -g '' elf
WORKDIR /opt/app/
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/accommodation ./cmd/app

FROM alpine:3.14
LABEL language="golang"
COPY .env .
COPY Makefile .
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=elf:1000 /go/bin/accommodation /accommodation
USER elf
ENTRYPOINT ["./accommodation"]