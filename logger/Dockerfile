FROM python:3.9.6-alpine

# Keeps Python from generating .pyc files in the container
ENV PYTHONDONTWRITEBYTECODE 1
# Turns off buffering for easier container logging
ENV PYTHONUNBUFFERED 1

RUN apk update && \
    apk add curl musl-dev openssl-dev libffi-dev && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python3 -

ENV PATH="${PATH}:/root/.poetry/bin"

WORKDIR /app

COPY . .

RUN poetry config virtualenvs.create false && \
    poetry install --no-dev

EXPOSE 8000

CMD [ "uvicorn", "--port", "8080", "--host", "0.0.0.0", "--app-dir", "src", "app:app"]
