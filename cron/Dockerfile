FROM python:3.9-alpine

WORKDIR /app

COPY ./program/* /app/

RUN apk update && apk add --no-cache cron bash
RUN pip install -r requirements.txt

COPY cronjob /etc/cron.d/cronjob

RUN chmod 0644 /etc/cron.d/cronjob
RUN crontab /etc/cron.d/cronjob

CMD cron && tail -f /var/log/cron.log
