FROM golang:latest

WORKDIR /APP

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN GOOS=linux go build -o app1 

EXPOSE 4545

CMD [ "/app1" ]

