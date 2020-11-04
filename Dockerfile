FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go build

EXPOSE 8000

CMD ["/go/src/app/unictelezioni"]