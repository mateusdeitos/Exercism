FROM  --platform=linux/amd64 ubuntu:18.04
ARG TOKEN

RUN apt-get update && \
    apt-get -y dist-upgrade && \
    apt-get install -y \
    wget \
    tree \
    vim && \
    apt-get clean

RUN wget https://github.com/exercism/cli/releases/download/v3.3.0/exercism-3.3.0-linux-arm64.tar.gz && \
    tar xzf exercism-3.3.0-linux-arm64.tar.gz && \
    mv exercism /usr/local/bin/

RUN /usr/local/bin/exercism configure --token=$TOKEN

RUN mkdir /workspace
WORKDIR /workspace
