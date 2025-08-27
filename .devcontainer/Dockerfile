FROM golang:1.23

# Diretório de trabalho
WORKDIR /workspace

# Dependências úteis
RUN apt-get update && apt-get install -y \
    git zsh curl vim \
    && rm -rf /var/lib/apt/lists/*

# Docker Para Debian/Ubuntu
RUN apt-get update && \
    apt-get install -y docker.io && \
    rm -rf /var/lib/apt/lists/*


# Instalar Oh My Zsh
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended \
    && chsh -s $(which zsh)

# Adicionar o GOPATH/bin ao PATH
ENV PATH=$PATH:/go/bin

# Expor porta do Go
EXPOSE 8080

# Mantém container ativo
CMD ["sleep", "infinity"]
ENTRYPOINT ["tail", "-f", "/dev/null"]
