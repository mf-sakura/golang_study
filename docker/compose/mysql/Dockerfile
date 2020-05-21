FROM mysql:5.7

# RUN usermod -u 201 -g 100 mysql

ADD ./conf.d/my.cnf /etc/mysql/my.cnf
ADD ./docker-entrypoint-initdb.d /docker-entrypoint-initdb.d/

RUN chmod 644 /etc/mysql/my.cnf
