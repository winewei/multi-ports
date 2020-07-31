FROM alpine

WORKDIR /srv

COPY main main

EXPOSE 8000

CMD ["/srv/main"]
