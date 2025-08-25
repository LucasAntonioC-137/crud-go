# Imagem oficial do Go (versão estável)
FROM golang:1.22-bullseye

# Definir diretório de trabalho dentro do container
WORKDIR /workspace

# Instalar dependências extras (útil para desenvolvimento)
RUN apt-get update && apt-get install -y \
    git \
    zsh \
    curl \
    vim \
    && rm -rf /var/lib/apt/lists/*

# Instalar Oh My Zsh (opcional, para um terminal mais agradável)
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended \
    && chsh -s $(which zsh)

# Adicionar o GOPATH/bin ao PATH
ENV PATH=$PATH:/go/bin

# Expor a porta que o Go pode usar (ajuda em servidores web locais)
EXPOSE 8080

# Comando padrão: abre o zsh
CMD [ "zsh" ]
