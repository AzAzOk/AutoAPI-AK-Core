**AutoAPI-AK-Core** is a bold project to transform scattered trial access, promo credits, and temporary API limits into one intelligent system for developers. We are building a world where every key, every free request, and every temporary access window stops being lost in tab chaos and becomes part of a single, coordinated engine. AutoAPI-AK-Core helps teams connect external services faster, route requests smarter, and turn short-lived opportunities into real product momentum.

## Why It Matters

Today, many services offer free credits, test environments, and limited-time access, but almost nobody turns them into something centralized, elegant, and genuinely powerful. AutoAPI-AK-Core turns that fragmentation into a strategic advantage.

## The Vision

- Unified control over free and temporary API access.
- Smarter integration with external services.
- Automated routing, fallback, and key management.
- Bright. Elegantly. Inevitable.


# AutoAPI-AK-Core

> Единый прокси-оркестратор для управления пулом аккаунтов AI-сервисов с автоматической ротацией и сохранением контекста.

## Что это

Позволяет использовать несколько аккаунтов одного сервиса (OpenAI, Claude, Gemini и т.д.) как единое целое:
- Автоматически переключается между аккаунтами при исчерпании лимитов
- Сохраняет контекст диалога при смене аккаунта
- Предоставляет единый API наружу

## Сервисы

| Сервис | Язык | Ответственный | Описание |
|---|---|---|---|
| `gateway` | Node.js | Друг | Единая точка входа, проксирование |
| `account-vault` | Node.js | Друг | Хранение аккаунтов и ключей |
| `auth` | Node.js | Друг | JWT авторизация, роли |
| `router` | Go | Ты | Логика выбора аккаунта, ротация |
| `context-bridge` | Go | Ты | Хранение и восстановление контекста |

## Быстрый старт

```bash
# 1. Клонировать репо
git clone https://github.com/yourorg/autoapi-ak-core.git
cd autoapi-ak-core

# 2. Настроить env
make setup

# 3. Запустить
make up
```

## Команды

```bash
make up          # Поднять все сервисы
make down        # Остановить
make test        # Все тесты
make lint        # Линтинг
make build-go    # Собрать Go сервисы
make proto       # Генерация из .proto
```

## Структура

```
services/
  gateway/         # Node.js — входная точка
  account-vault/   # Node.js — хранилище аккаунтов
  auth/            # Node.js — авторизация
  router/          # Go    — роутер запросов
  context-bridge/  # Go    — мост контекста
packages/
  adapters/        # Адаптеры под конкретные AI-сервисы
  crypto/          # Утилиты шифрования
  types/           # Общие типы
docs/
  adr/             # Architecture Decision Records
  runbooks/        # Инструкции по восстановлению
infra/
  docker/          # Docker конфиги
  k8s/             # Kubernetes манифесты
```

## Безопасность

- Все ключи шифруются AES-256-GCM перед сохранением
- `.env` файлы никогда не коммитятся (см. `.gitignore`)
- Секреты ротируются каждые 30-90 дней
