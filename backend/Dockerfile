FROM golang:1.13.6-alpine3.10 AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git

ARG API_SRC
WORKDIR $API_SRC
COPY . $API_SRC

EXPOSE 8080
