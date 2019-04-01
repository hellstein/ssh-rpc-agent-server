FROM alpine

RUN apk update && \
apk upgrade && \
apk add bash openssh vim sshpass

RUN mkdir /credential
COPY ssh-rpc-agent-server ./


CMD ["/bin/bash", "-c", "./ssh-rpc-agent-server"]
