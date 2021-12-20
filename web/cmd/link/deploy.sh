#!/bin/sh


pwd

rm -rf /root/app/front-end/*
tar -xf /root/app/package/package-front.tgz -C /root/app/front-end

if [ /root/app/backend/pidfile.txt ]; then
    kill  `cat /root/app/backend/pidfile.txt`
fi

tar -xvf package/package-backend.tgz -C /root/app/backend

cd /root/app/backend

nohup ./link 2>&1 > log.txt &
echo $! > pidfile.txt
