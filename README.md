# Komodo Telegram Alerter

A lightweight webhook receiver that forwards Komodo monitoring alerts to Telegram.

## Quick Start

### Using Docker

With config file:

```bash
docker run -d \
  -p 8964:8964 \
  -v /path/to/config.yaml:/app/config.yaml \
  ronmi/kta:latest serve -c /app/config.yaml
```

With environment variables:

```bash
docker run -d \
  -p 8964:8964 \
  --env-file /path/to/.env \
  ronmi/kta:latest serve
```

Docker image: [ronmi/kta](https://hub.docker.com/repository/docker/ronmi/kta/general)

### Using Docker Compose

With config file:

```yaml
services:
  kta:
    image: ronmi/kta:latest
    ports:
      - "8964:8964"
    volumes:
      - ./config.yaml:/app/config.yaml
    command: serve -c /app/config.yaml
```

With environment variables:

```yaml
services:
  kta:
    image: ronmi/kta:latest
    ports:
      - "8964:8964"
    env_file:
      - .env
    command: serve
```

### Direct Execution

```bash
kta serve -c /path/to/config.yaml
```

## Configuration

### Using Config File

Create a YAML configuration file (see `example.komodo-tg-alerter.yaml`):

```yaml
web:
  bind: ":8964"
log:
  level: info
telegram:
  token: your_telegram_bot_token
  chat: your_chat_id
```

### Using Environment Variables

Alternatively, use environment variables (see `example.komodo-tg-alerter.env`):

```bash
KTA_WEB_BIND=:8964
KTA_LOG_LEVEL=info
KTA_TELEGRAM_TOKEN=your_telegram_bot_token
KTA_TELEGRAM_CHAT=your_chat_id
```

## Building from Source

Requirements: Go 1.21+

```bash
git clone https://github.com/raohwork/komodo-tg-alerter.git
cd komodo-tg-alerter
go build -o kta
```

## License

GPL-3.0
