FROM golang:1.15.3
RUN apt-get update

WORKDIR /app
COPY src/* /app
RUN go build main.go

EXPOSE 80/tcp
CMD ["./main"]