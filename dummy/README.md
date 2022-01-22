# Dummy

Dummy contains two main features:
1. Generate a whole lot of dummy data
2. Insert the dummy data into the InfluxDB database

## Usage

Please use Python 3.10, you can download [here](https://www.python.org/downloads/release/python-3101/) make sure you choose as you machine CPU Arch (usually 64-bit for most laptop, except MacBook which use M1 Processor that use ARM arch) and the tutorial [Install Python](https://gadiskoding.my.id/install-python-di-komputer-dan-android.html).

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

| Env variable | How to get |
|--|--|
| INFLUX_TOKEN | This influx authorization token, usually come with user account make it in GUI influx |
| INFLUX_HOST | This you influxdb server address to connect with the services |
| INFLUX_ORG | This you influx db org |

Just put the value like this, for example `INFLUX_HOST=localhost:8086`.

If you use docker compose with same as influxdb, just make the `INFLUX_ORG` match with `DOCKER_INFLUXDB_INIT_ORG` environment variable in variable and `INFLUX_HOST` the name of container with it port, for example `influx_services:8086` .

To run the insertion script:

```sh
pipenv run python3 inserter.py
```

Enjoy!
