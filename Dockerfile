FROM gcr.io/distroless/base

COPY ./build/linux/gcs-copy /usr/bin/gcs-copy

CMD ["gcs-copy", "version"]
