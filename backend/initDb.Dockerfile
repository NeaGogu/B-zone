# Description: This file is used to initialize the database with the data from the test_binding folder
# Author: Gogu
# It's not meant to be used freely, just a one time thing ;)

FROM mongo:latest
# COPY /home/gogu/Documents/test_binding /data/db
# RUN mkdir -p /data/db2 \
#     && echo "dbpath = /data/db2" > /etc/mongodb.conf \
#     && chown -R mongodb:mongodb /data/db2 \
#     && chmod -R 777 /data/db2


ADD ./test_binding /data/db
RUN chown -R mongodb:mongodb /data/db

USER mongodb
CMD ["mongod", "--dbpath=/data/db"]
