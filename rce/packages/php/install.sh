#!/usr/bin/env bash

apt-get install -y apt-transport-https lsb-release ca-certificates curl

curl -O /etc/apt/trusted.gpg.d/php.gpg https://packages.sury.org/php/apt.gpg

sh -c 'echo "deb https://packages.sury.org/php/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/php.list'

apt-get update -y

apt-get install -y php8.1
