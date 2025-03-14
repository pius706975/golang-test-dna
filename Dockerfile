FROM golang:1.24.1-alpine3.21

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -v -o /app/go-test

EXPOSE 5000

ENTRYPOINT [ "/app/go-test" ]
CMD [ "serve" ]