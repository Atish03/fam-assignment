FROM python:3.9-alpine

WORKDIR /app

COPY ./program/* .
COPY ./requirements.txt .

RUN pip install -r requirements.txt

COPY cronjob /etc/cron/crontab

RUN crontab /etc/cron/crontab

CMD ["crond", "-f", "-l", "0"]
