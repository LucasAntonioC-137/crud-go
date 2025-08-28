FROM golang:1.23

WORKDIR /workspace

RUN apt-get update && apt-get install -y \
    git zsh curl vim docker.io \
    && rm -rf /var/lib/apt/lists/*

RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended \
    && chsh -s $(which zsh)

ENV PATH=$PATH:/go/bin

EXPOSE 8080
