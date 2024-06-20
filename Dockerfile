FROM golang:1.21 AS builder
RUN apt-get update -y
RUN apt-get install git -y
WORKDIR /callisto
COPY . ./
RUN make build

FROM alpine:latest
WORKDIR /callisto
COPY --from=builder /callisto /usr/bin/callisto
CMD [ "callisto" ]