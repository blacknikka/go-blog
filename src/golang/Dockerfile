FROM golang:1.14

WORKDIR /go/src/github.com/blacknikka/go-blog

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.7.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv ./migrate.linux-amd64 /usr/bin/

ENV PATH $PATH:/go/bin

# CMD ["go", "run", "./main.go"]
