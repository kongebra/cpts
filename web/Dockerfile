FROM nginx

RUN adduser --disabled-password --gecos '' web
USER web

WORKDIR ./web

COPY . /usr/share/nginx/html