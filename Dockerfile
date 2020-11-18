FROM golang:1.15-alpine as builder
RUN mkdir /build
COPY src /build
WORKDIR /build
RUN go build -o main .

FROM alpine
COPY --from=builder /build/main /app/
#COPY .env /
WORKDIR /app 

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

CMD ["./main"]
