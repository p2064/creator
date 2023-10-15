FROM golang:1.20


ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go build -o /app/creator

EXPOSE 9000


# Run
CMD ["/app/creator"]