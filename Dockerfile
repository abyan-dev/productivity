FROM alpine:3.20

WORKDIR /app

COPY ./productivity /app

CMD ["/app/productivity"]