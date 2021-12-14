FROM golang:1.17.2-alpine3.14
WORKDIR /app
RUN [ "apk", "add", "curl" ]
RUN [ "apk", "add", "make" ]
COPY go.mod .
COPY go.sum .
COPY makefile .
RUN [ "make", "prepare" ]
COPY . .
CMD [ "make", "run" ]