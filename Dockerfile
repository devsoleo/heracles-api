FROM alpine:latest

WORKDIR /app

COPY myapp .
COPY database.db .

RUN chmod +x /app/myapp

EXPOSE 8080

CMD ["/app/myapp"]