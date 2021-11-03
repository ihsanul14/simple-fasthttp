FROM golang:1.14.3

COPY . /app
WORKDIR /app

RUN go mod vendor
RUN go build

ENTRYPOINT [ "./simple-fasthttp" ]