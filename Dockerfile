FROM golang:1.15-alpine as builder
RUN mkdir /build
COPY src/* /build
WORKDIR /build
RUN go build -o main .

FROM alpine
COPY --from=builder /build/main /app/
WORKDIR /app 

EXPOSE 10000/tcp
CMD ["./main"]
