# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY gameApiApp /app

CMD [ "/app/gameApiApp" ]