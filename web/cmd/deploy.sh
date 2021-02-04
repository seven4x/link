#!/bin/sh


rm -rf /root/app/front-end/*
tar -xf /root/app/package-front.tgz -C /root/app/front-end

cd /root/app

if [ --f pidfile.txt ]; then
    kill  `cat pidfile.txt`
fi


nohup ./link 2>&1 > log.txt & echo $! > pidfile.txt

