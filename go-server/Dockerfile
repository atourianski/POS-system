FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get github.com/go-sql-driver/mysql

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
