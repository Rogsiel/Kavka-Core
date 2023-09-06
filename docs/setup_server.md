# Setup Server

## Configs

```yml
APP:
  NAME: "kavka"
  HTTP:
    HOST: "127.0.0.1"
    PORT: 8000
    ADDRESS: "127.0.0.1:8000"
  FIBER:
    SERVER_HEADER: "Fiber"
    PREFORK: false
    CORS:
      ALLOW_ORIGINS: "*"
      ALLOW_CREDENTIALS: true
MONGO:
  HOST: "127.0.0.1"
  USERNAME: "mongo"
  PASSWORD: "mongo"
  PORT: 27017
  DB_NAME: "kavka"
REDIS:
  HOST: "127.0.0.1"
  USERNAME: "redis"
  PASSWORD: "redis"
  PORT: 6379
```