FROM golang:1.21-alpine AS builder

RUN apk add build-base git linux-headers
WORKDIR /work
COPY . ./
RUN make build

FROM alpine:latest
WORKDIR /work
COPY --from=builder /work/build/callisto /usr/bin/callisto