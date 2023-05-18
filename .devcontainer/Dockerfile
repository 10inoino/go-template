FROM golang:1.19.9

# llのエイリアスを指定
RUN echo "alias ll='ls -l'" >> /root/.bashrc

# プロンプト設定(ref. https://zenn.dev/daifukuninja/articles/9e903808f4147d)
RUN apt update && apt install git vim curl -y && apt clean

# nodejs, yarnのインストール
RUN curl https://deb.nodesource.com/setup_20.x | bash
RUN curl https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list

RUN apt-get update && apt-get install -y nodejs yarn apt-transport-https ca-certificates gnupg

# プロンプト設定
RUN echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc

RUN curl -O https://raw.githubusercontent.com/git/git/master/contrib/completion/git-prompt.sh
RUN curl -O https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash
RUN chmod a+x git*.*
RUN ls -l $PWD/git*.* | awk '{print "source "$9}' >> ~/.bashrc

RUN echo "GIT_PS1_SHOWDIRTYSTATE=true" >> ~/.bashrc
RUN echo "GIT_PS1_SHOWUNTRACKEDFILES=true" >> ~/.bashrc
RUN echo "GIT_PS1_SHOWUPSTREAM=auto" >> ~/.bashrc

RUN echo 'export PS1="\[\033[01;32m\]\u@\h\[\033[01;33m\] \w \[\033[01;31m\]\$(__git_ps1 \"(%s)\") \\n\[\033[01;34m\]\\$ \[\033[00m\]"' >> ~/.bashrc

# 必要なGoのパッケージの追加（gopls, sql-migrate, sql-boiler）
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
RUN apt-get update && apt-get install -y google-cloud-cli mariadb-client

# planetscale CLIのインストール
RUN curl -OL https://github.com/planetscale/cli/releases/download/v0.142.0/pscale_0.142.0_linux_arm64.deb
RUN dpkg -i pscale_0.142.0_linux_arm64.deb
