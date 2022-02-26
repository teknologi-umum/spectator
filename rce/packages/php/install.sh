#!/usr/bin/env bash

apt install -y lsb-release ca-certificates apt-transport-https software-properties-common gnupg2 wget

echo "deb https://packages.sury.org/php/ $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/sury-php.list

wget -qO - https://packages.sury.org/php/apt.gpg | sudo apt-key add -

apt update

apt install php8.1
