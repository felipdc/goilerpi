FROM golang:1.19-alpine

RUN apk update && apk add git

COPY . /app
WORKDIR /app

RUN go get -d -v

RUN go build

EXPOSE 4000

CMD [ "./goilerpi" ]