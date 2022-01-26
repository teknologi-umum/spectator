FROM python:3.10.2-alpine3.15

WORKDIR /usr/src/app

COPY . .
RUN apk add build-base \
    musl-dev \
    libffi-dev \
    python3-dev \
    openssl-dev
RUN pip install pipenv
RUN pipenv install

RUN [ "pipenv", "run", "python3", "generate.py" ]
RUN [ "pipenv", "run", "python3", "inserter.py" ]