FROM postgres:15.2

ARG XVIEWER_VERSION=2.1.0
ARG XVIEWER_RELEASE_TYPE=SNAPSHOT

#Copy required initialization files
COPY init/*.sql /docker-entrypoint-initdb.d/