FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# ビルドしないで go run で直接実行
CMD ["go", "run", "main.go"]