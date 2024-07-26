FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY build/uttt /app/uttt
COPY static /app/static

ENTRYPOINT [ "/app/uttt" ]
