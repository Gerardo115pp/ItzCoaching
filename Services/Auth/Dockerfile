FROM golang:1.19.0-alpine3.15 AS builder
RUN apk add --no-cache git gcc musl-dev
WORKDIR /go/src/auth_service
COPY . .
RUN go mod download
RUN go build -o /go/bin/auth_service

FROM alpine:3.15.0
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/auth_service /auth_service
RUN mkdir /operation_data
ENTRYPOINT /auth_service
LABEL Name=AuthService Version=1.0
EXPOSE 4052
CMD ["auth_service"]