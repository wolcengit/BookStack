#!/bin/sh

if [ ! -f "/bookstack/conf/app.conf" ] ; then
  cd /bookstack/conf
  cp app.conf.example app.conf
  cp oauth.conf.example oauth.conf
  cp oss.conf.example oss.conf
  sed -i 's/runmode = dev/runmode = prod/g' app.conf
fi

cd /bookstack
if [ ! -d "logs" ] ; then
  mkdir -p logs
fi
if [ ! -d "store" ] ; then
  mkdir -p store
fi
if [ ! -d "uploads" ] ; then
  mkdir -p uploads
fi

/bookstack/BookStack install
/bookstack/BookStack