FROM centos:7

CMD ["gcs-copy", "version"]

COPY ./build/linux/gcs-copy /usr/bin/gcs-copy
