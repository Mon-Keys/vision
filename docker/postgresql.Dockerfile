FROM ubuntu:20.04

RUN apt-get -y update && apt-get install -y tzdata
ENV TZ=Russia/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get -y update && apt-get install -y postgresql
USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --command "ALTER USER postgres WITH PASSWORD 'password';" &&\
    createdb -O postgres forum &&\
    /etc/init.d/postgresql stop

EXPOSE 5432
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]