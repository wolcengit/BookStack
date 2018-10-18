# image wolcen/bookstack:1.4-Z16
FROM wolcen/calibre:1.1
LABEL author="wolcen@msn.com"


COPY lib/time/zoneinfo.zip /usr/local/go/lib/time/
ADD conf/*.conf.example /bookstack/conf/
ADD dictionary /bookstack/dictionary/
ADD static /bookstack/static/
ADD *.md /bookstack/
ADD favicon.ico /bookstack/
ADD start.sh /bookstack/
ADD views /bookstack/views/
ADD BookStack /bookstack/

WORKDIR /bookstack
RUN chmod +x /bookstack/BookStack
RUN chmod +x /bookstack/start.sh

EXPOSE 8181
ENTRYPOINT ["/bookstack/start.sh"]
