# Dummy

Dummy contains two main features:
1. Generate a whole lot of dummy data
2. Insert the dummy data into the InfluxDB database

## Usage

Please use Python 3.10.

For Windows, Download the [installer here](https://www.python.org/ftp/python/3.10.0/python-3.10.0-amd64.exe),
then follow the instruction (just press next until the installation is finished)

For Mac, Download the [installer here](https://www.python.org/ftp/python/3.10.0/python-3.10.0post2-macos11.pkg),
then follow the instruction there.

For Linux, use package manager or build from source:
- Arch (latest): `pacman -S python3`
- Alpine (latest): `apk add python3=3.10.1-r0`
- Ubuntu: `apt install python3`
- Other - by building from source:
  - Install Make, GCC, and development libraries:
    - Debian-based: `apt-get install build-essential make gcc`
    - RHEL-based: `yum groupinstall 'Development Tools'`
  - Download Python source code: `curl -O https://www.python.org/ftp/python/3.10.7/Python-3.10.7.tgz`
  - Extract the source code : `tar -zxf Python-3.10.7.tgz`
  - Change directory to `Python-3.10.8` by doing `cd Python-3.10.7`.
  - Configure installation: `./configure --prefix=/usr`
  - Compile and install: `make -j $(nproc) && sudo make install -j $(nproc)`

Setup first.

```sh
# Install pipenv first:
pip install --user pipenv

# Install dependencies
pipenv install
```

To generate user personal info (including their session) and the events, do:
```sh
pipenv run python3 generate.py
```

Then, we shall insert the generated data into the InfluxDB database. But first,
we need to put some data into the environment variable. The easiest way to do
so is to copy .env.example and rename it into .env file.

| Environment variable key | Description                                                                                   |
| ------------------------ | --------------------------------------------------------------------------------------------- |
| INFLUX_TOKEN             | This influx authorization token, usually come with user account make it in InfluxDB dashboard |
| INFLUX_HOST              | This is your InfluxDB server address                                                          |
| INFLUX_ORG               | This is your InfluxDB user organization's name                                                |

Just put the value like this, for example `INFLUX_HOST=localhost:8086` at you `.env` files.

If you use docker compose with same as influxdb, just make the `INFLUX_ORG` match with `DOCKER_INFLUXDB_INIT_ORG`
environment variable in variable and `INFLUX_HOST` the name of container with it port, for example `influx_services:8086`.

To run the insertion script:

```sh
pipenv run python3 inserter.py
```
