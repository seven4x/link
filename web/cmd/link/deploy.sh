#!/bin/sh


pwd
cd /root/app/butterfly/deploy_tmp
# restart bd
ps -ef | grep link | grep -v grep | cut -c 9-15  | xargs kill
rm -rf ../bd/*
tar -xf package-backend.tgz  -C ../bd

# replace front
rm -rf ../www/*
tar -xf package-front.tgz -C ../www

#start bd
cd ../bd
nohup ./link 2>&1 > log.txt &
echo $! > pidfile.txt

echo "success"