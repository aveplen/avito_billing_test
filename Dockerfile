FROM golang:1.17

WORKDIR $GOPATH/src/aveplen/avito_billing_test
COPY . .

RUN apt update
RUN apt -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o avito_billing ./cmd/main.go

CMD ["./avito_billing"]