import Fastify from 'fastify'

const app = Fastify({ logger: true })
const PORT = Number(process.env.ACCOUNT_VAULT_PORT) || 3002

app.get('/healthz', async () => ({ status: 'ok' }))
app.get('/readyz', async () => ({ status: 'ready' }))

// Routes (заглушки — реализовать)
// POST   /v1/accounts        — добавить аккаунт
// GET    /v1/accounts        — список аккаунтов пользователя
// DELETE /v1/accounts/:id    — удалить аккаунт
// GET    /v1/accounts/available/:provider — получить доступный аккаунт

app.listen({ port: PORT, host: '0.0.0.0' }, (err) => {
  if (err) { app.log.error(err); process.exit(1) }
})
