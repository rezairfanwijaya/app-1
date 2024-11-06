FROM golang:latest

WORKDIR /APP

COPY . .

RUN go mod tidy

RUN GOOS=linux go build -o app1 

EXPOSE 4545

ENTRYPOINT [ "./app1" ]

