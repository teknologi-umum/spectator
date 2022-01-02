# Dummy

Dummy contains two main features:
1. Generate a whole lot of dummy data
2. Insert the dummy data into the InfluxDB database

## Usage

Please use Python 3.10.

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
```sh
pipenv run python3 inserter.py
```

Enjoy!
