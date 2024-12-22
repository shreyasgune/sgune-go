# STAGE 1
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o personApi .

# STAGE 2
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/personApi /personApi

EXPOSE 2112

CMD ["/personApi"]