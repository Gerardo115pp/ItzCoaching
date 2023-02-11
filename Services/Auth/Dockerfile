FROM golang:1.20.0-alpine3.16 AS builder
RUN apk add --no-cache git gcc musl-dev
WORKDIR /go/src/auth_service
COPY . .
RUN go mod download
RUN go build -o /go/bin/auth_service

FROM alpine:3.16.0
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/auth_service /auth_service
RUN mkdir /operation_data
ENTRYPOINT /auth_service
LABEL Name=ItzAuthService Version=1.0
EXPOSE 4152
CMD ["auth_service"]