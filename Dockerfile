FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go mod tidy

RUN go build -o /phone-number-lookup

EXPOSE 8080

CMD [ "/phone-number-lookup" ]