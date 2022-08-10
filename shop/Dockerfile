FROM python:3.8

WORKDIR /shop

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

RUN apt-get update && \
    apt-get install -y wget lsb-release && \
    echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list && \
    (wget --no-check-certificate --quiet -O - https://apt.postgresql.org/pub/repos/apt/ACCC4CF8.asc | apt-key add -)

RUN apt-get update && \
    apt-get install -y postgresql-client-9.6 gettext cron && \
    apt-get clean

RUN pip install --no-cache-dir --upgrade pip setuptools
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

RUN chmod 0700 entrypoint.sh
CMD ["/bin/bash", "entrypoint.sh"]
