# image wolcen/bookstack:1.4-Z4
FROM wolcen/calibre
LABEL author="wolcen@msn.com"


ADD conf/*.conf.example /bookstack/conf/
ADD dictionary /bookstack/dictionary/
ADD static /bookstack/static/
ADD views /bookstack/views/
ADD *.md /bookstack/
ADD favicon.ico /bookstack/
ADD BookStack /bookstack/
ADD start.sh /bookstack/
COPY lib/time/zoneinfo.zip /usr/local/go/lib/time/

WORKDIR /bookstack
RUN chmod +x /bookstack/BookStack
RUN chmod +x /bookstack/start.sh

EXPOSE 8181
ENTRYPOINT ["/bookstack/start.sh"]
