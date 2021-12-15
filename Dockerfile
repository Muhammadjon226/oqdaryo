FROM alpine:edge

RUN mkdir /app
COPY oqdaryo /app
WORKDIR /app

CMD ./oqdaryo
