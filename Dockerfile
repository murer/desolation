FROM murer/hexblade:edge

USER root

RUN apt-get -y update
RUN apt-get -y install git wget xauth net-tools dnsutils inetutils-syslogd vim zbar-tools xdotool curl imagemagick
RUN apt-get -y install openssh-server
RUN apt-get -y install build-essential

RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | sudo apt-key add - && \
    echo 'deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main' > /etc/apt/sources.list.d/google-chrome.list && \
    apt-get -y update && \
    apt-get -y install google-chrome-stable

USER hexblade
WORKDIR /home/hexblade
RUN mkdir -p .vscode openerssh go/packages

RUN wget --progress=dot -e dotbytes=1M -c \
        'https://go.microsoft.com/fwlink/?LinkID=620884' \
        -O vscode.tar.gz && \
    tar xzf vscode.tar.gz && \
    rm vscode.tar.gz && \
    mv VSCode-linux-x64 vscode

RUN cd go && \
    wget --progress=dot -e dotbytes=1M -c \
        'https://golang.org/dl/go1.14.10.linux-amd64.tar.gz' \
        -O go.tar.gz && \
    tar xzf go.tar.gz && \
    rm go.tar.gz

ENV GOROOT "$HOME/go/go"
ENV GOPATH "$HOME/go/packages"
ENV PATH "$HOME/vscode/bin:$GOROOT/bin:$PATH"

RUN go get -u github.com/ramya-rao-a/go-outline
RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get -u github.com/sqs/goreturns
RUN go get -u github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get -u github.com/stamblerre/gocode
RUN go get -u github.com/rogpeppe/godef
RUN go get -u github.com/acroca/go-symbols
RUN go get -u golang.org/x/tools/gopls
RUN go get -u golang.org/x/lint/golint

COPY docker /opt/openerssh/docker
