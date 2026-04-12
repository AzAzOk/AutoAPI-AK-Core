# Runbook: Восстановление сервисов

## Диагностика

```bash
# Проверить статус всех контейнеров
docker compose ps

# Посмотреть логи конкретного сервиса
docker compose logs -f router
docker compose logs -f context-bridge

# Проверить healthcheck
curl http://localhost:8080/healthz   # router
curl http://localhost:8081/healthz   # context-bridge
curl http://localhost:3000/healthz   # gateway
curl http://localhost:3001/healthz   # auth
curl http://localhost:3002/healthz   # account-vault
```

## Рестарт сервиса

```bash
docker compose restart router
docker compose restart context-bridge
```

## Полный перезапуск

```bash
make down && make up
```

## Проблема: сервис не стартует

1. Проверить `.env` — все переменные заполнены?
2. Проверить PostgreSQL: `docker compose logs postgres`
3. Проверить Redis: `docker compose logs redis`
4. Проверить порты: `lsof -i :8080`

## Откат деплоя

```bash
git log --oneline -10          # найти хороший коммит
git checkout <commit-hash>
make up
```
