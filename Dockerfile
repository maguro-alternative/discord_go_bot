FROM golang:1.20.4-bullseye AS builder

RUN apt-get -y update && apt-get -y install locales && apt-get -y upgrade && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm

# ./root/src ディレクトリを作成 ホームのファイルをコピーして、移動
RUN mkdir -p /root/src
COPY . /root/src
WORKDIR /root/src

# Docker内で扱うffmpegをインストール
RUN apt-get install -y ffmpeg

RUN go mod download

RUN go build -o /app/main ./main.go

# ポート8080を外部に公開
EXPOSE 8080

# アプリケーションを実行
CMD ["/app/main"]

# Runner用の新しいステージを開始
#FROM debian:latest AS runner

#RUN apt-get update && apt-get install -y locales && \
#    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
#ENV LANG ja_JP.UTF-8
#ENV LANGUAGE ja_JP:ja
#ENV LC_ALL ja_JP.UTF-8
#ENV TZ JST-9
#ENV TERM xterm

# ビルダーステージからバイナリをコピー
#COPY --from=builder /app/main /app/main