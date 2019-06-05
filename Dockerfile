# This file is a template, and might need editing before it works on your project.
FROM golang:1.12.1-alpine3.9 AS builder

WORKDIR /build
WORKDIR /src

COPY . .
# RUN go-wrapper download
ENV GO111MODULES=on
RUN go build -v -mod vendor -o /build/fizzbuzz

FROM alpine:3.9

# We'll likely need to add SSL root certificates
RUN apk --no-cache add ca-certificates

WORKDIR /usr/local/bin

ARG main_folder
COPY --from=builder /build/fizzbuzz .
CMD ["./fizzbuzz"]
