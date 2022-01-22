# Dummy

Dummy contains two main features:
1. Generate a whole lot of dummy data
2. Insert the dummy data into the InfluxDB database

## Usage

Please use Python 3.10.

For windows, Download the [installer here](https://www.python.org/ftp/python/3.10.0/python-3.10.0-amd64.exe), then follow the instruction (just press next until the installation is finished)

For Mac, Download the [installer here](https://www.python.org/ftp/python/3.10.0/python-3.10.0post2-macos11.pkg), then follow the instruction there.

For Linux, many linux distribution some are require to build from scratch (e.g. Ubuntu).

- Arch (latest): `pacman -S python3`

- Alpine (latest): `apk add python3=3.10.1-r0`

- Ubuntu (for ubuntu you need to make it from scratch) :

  - Install ubuntu dev tools : `sudo apt install build-essential`
  - Download the python source code : `curl -O https://www.python.org/ftp/python/3.10.0/Python-3.10.0.tgz`
  - Extract the source code : `sudo tar -xvzf Python-3.10.0.tgz`
  - Go inside `Python-3.10.0`, `cd Python-3.10.0`.
  - Then run this to start compile: `./configure --prefix=/usr && make && make install`

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

|  Environment variable key | Description |
|--|--|
| INFLUX_TOKEN | This influx authorization token, usually come with user account make it in InfluxDB dashboard |
| INFLUX_HOST | This is your InfluxDB server address |
| INFLUX_ORG |  This is your InfluxDB user organization's name  |

Just put the value like this, for example `INFLUX_HOST=localhost:8086` at you `.env` files.

If you use docker compose with same as influxdb, just make the `INFLUX_ORG` match with `DOCKER_INFLUXDB_INIT_ORG` environment variable in variable and `INFLUX_HOST` the name of container with it port, for example `influx_services:8086` .

To run the insertion script:

```sh
pipenv run python3 inserter.py
```

Enjoy!
