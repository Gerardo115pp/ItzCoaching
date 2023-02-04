FROM golang:1.19.0-alpine3.15 AS builder
RUN apk add --no-cache git gcc musl-dev
WORKDIR /go/src/customers_service
COPY . .
RUN go mod download
RUN go build -o /go/bin/customers_service

FROM alpine:3.15.0
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/customers_service /customers_service
RUN mkdir /operation_data
ENTRYPOINT /customers_service
LABEL Name=ItzCustomerService Version=1.0
EXPOSE 4151
CMD ["customers_service"]