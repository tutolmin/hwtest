FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o hwtest hwtest.go
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/hwtest /app/
COPY --from=builder /build/hwtest.html /app/
COPY --from=builder /build/version /app/
WORKDIR /app
EXPOSE 8080
CMD ["./hwtest"]
