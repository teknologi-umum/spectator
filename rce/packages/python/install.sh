#!/usr/bin/env bash

VERSION="3.10.2"

cd /tmp
# Acquire the Python tar.gz file from the source repository
curl -O https://www.python.org/ftp/python/${VERSION}/Python-${VERSION}.tgz
tar -xvzf Python-${VERSION}.tgz
cd Python-${VERSION}
./configure \
  --prefix=/opt/python/${VERSION} \
  --enable-shared \
  --enable-ipv6 \
  LDFLAGS=-Wl,-rpath=/opt/python/${VERSION}/lib,--disable-new-dtags
make
make install

cd /tmp
sudo rm Python-${VERSION}.tgz
sudo rm -rf Python-${VERSION}
sudo rm get-pip.py
