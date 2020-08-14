FROM golang:1.14

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN  go build -o main cmd/proverb/main.go

CMD ["/app/main"]

