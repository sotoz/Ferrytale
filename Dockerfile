FROM golang:1.9

WORKDIR /go/src/app
COPY . .

ENV PORT 33333 
ENV DATABASE_URL root:root@tcp(database:3306)/ferrytale?parseTime=true

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 33333

CMD ["go-wrapper", "run"]