# image wolcen/calibre
FROM ubuntu:17.10
LABEL author="wolcen@msn.com"

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 安装calibre
RUN apt-get update -y
RUN apt-get upgrade -y 
RUN apt-get install -y apt-utils
RUN apt-get install -y locales
RUN apt-get install -y xdg-utils
RUN apt-get install -y xz-utils
RUN apt-get install -y python
RUN apt-get install -y chromium-browser
RUN apt-get install -y calibre
RUN apt-get install -y fonts-wqy-zenhei
RUN apt-get install -y fonts-wqy-microhei
RUN apt-get clean
RUN rm -rf /var/cache/apt/* /var/lib/apt/lists/*
