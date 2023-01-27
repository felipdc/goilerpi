FROM golang:1.19-alpine

RUN apk update && apk add git

COPY . /app
WORKDIR /app

RUN go get -d -v

# EXPOSE 8080

# CMD [ "ls /goilerpi/src" ]