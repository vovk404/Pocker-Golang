# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY sessionRedisServiceApp /app

CMD [ "/app/sessionRedisServiceApp" ]