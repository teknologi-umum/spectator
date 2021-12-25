FROM python:3.9.9-slim-buster

ENV PYTHONFAULTHANDLER=1 \
  PYTHONUNBUFFERED=1 \
  PYTHONHASHSEED=random \
  PIP_NO_CACHE_DIR=off \
  PIP_DISABLE_PIP_VERSION_CHECK=on \
  PIP_DEFAULT_TIMEOUT=100 \
  POETRY_VERSION=1.1.12

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y netcat gcc curl musl-dev openssl libffi-dev && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /app

COPY . .

RUN pip install "poetry==$POETRY_VERSION" && \
    poetry config virtualenvs.in-project true && \
    poetry install --no-dev

EXPOSE 8000

CMD [ "poetry", "run", "uvicorn", "--port", "8080", "--app-dir", "logger", "app:app" ]
