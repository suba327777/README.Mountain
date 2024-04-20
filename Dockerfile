FROM golang:1.21

WORKDIR /app

# ソースコードのコピー
COPY . .

# ビルド
RUN go build -o /main .

CMD [ "/main" ]