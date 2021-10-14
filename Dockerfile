FROM golang:1.14.3

COPY . /app
WORKDIR /app

ENV DB_HOST=192.168.100.59
ENV DB_NAME=mydb
ENV DB_USER=root
ENV DB_PASS=A123b456c
ENV DB_PORT=3306
ENV PORT=30001

RUN go mod vendor
RUN go build

ENTRYPOINT [ "./simple-fasthttp" ]