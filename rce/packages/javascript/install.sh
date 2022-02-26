#!/usr/bin/env bash

curl -O https://nodejs.org/dist/latest-v16.x/node-v16.14.0-linux-x64.tar.gz
tar -zxf node-v16.14.0-linux-x64.tar.gz
cd node-v16.14.0-linux-x64
mv -v bin/* /usr/bin/
mv -v include/* /usr/include/
mv -v lib/* /usr/lib/
mv -v share/doc/* /usr/share/doc/
mv -v share/man/* /usr/share/man/
mv -v share/* /usr/share/
cd ..
rm -rf node-v16.14.0-linux-x64 node-v16.14.0-linux-x64.tar.gz
