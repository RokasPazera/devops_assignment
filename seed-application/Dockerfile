FROM alpine:3.20
RUN apk add --no-cache bash curl ca-certificates
WORKDIR /app
COPY goals.mmd goals.mmv add_data.sh ./
RUN chmod +x add_data.sh
CMD ["./add_data.sh"]
