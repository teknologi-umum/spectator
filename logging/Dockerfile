FROM python:3.8.10-alpine

ENV PYTHONUNBUFFERED 1

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends netcat && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY . .

RUN pip install -r requirements.txt

COPY . ./

EXPOSE 8000

CMD [ "uvicorn", "--port", "8080", "--host", "0.0.0.0", "--app-dir", "src", "app:app"]
