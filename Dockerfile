# syntax=docker/dockerfile:1

# Stage 1 Build
FROM golang:1-alpine AS builder
RUN go version
COPY . /usr/src/goDataCollector/
WORKDIR /usr/src/goDataCollector/
RUN set -x && \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -v .

RUN chmod 755 /usr/src/goDataCollector/goDataCollector

# Stage 2 Final Image
FROM alpine:3
LABEL "com.deleteonerror.vendor"="Delete On Error"


# copy executable
COPY --from=builder /usr/src/goDataCollector/goDataCollector /usr/local/bin/
# copy empty configuration
COPY --from=builder /usr/src/goDataCollector/appData/config.json /etc/goDataCollector/
# create Data Directory
RUN mkdir -p /var/lib/goDataCollector/Log

CMD [ "goDataCollector", "-d" ]


