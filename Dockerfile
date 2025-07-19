FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD .env /app/.env
WORKDIR /app

ADD notely /usr/bin/notely

CMD ["notely"]
