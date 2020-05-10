FROM golang:1.14.1-alpine3.11
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o server .

EXPOSE 8080