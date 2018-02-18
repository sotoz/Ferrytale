FROM golang:1.10

WORKDIR /go/src/app
COPY . .

ENV APPLICATION_HOST 127.0.0.1 
ENV APPLICATION_PORT 33333 
ENV DATABASE_URL root:root@tcp(127.0.0.1:3306)/ferrytale?parseTime=true

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 33333

CMD ["go-wrapper", "run"]