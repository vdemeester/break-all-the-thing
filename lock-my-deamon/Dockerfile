FROM alpine

RUN apk update && apk add sshfs

RUN mkdir -p /run/docker/plugins /mnt/state /mnt/volumes

COPY lock-my-daemon lock-my-daemon

CMD ["lock-my-daemon"]
