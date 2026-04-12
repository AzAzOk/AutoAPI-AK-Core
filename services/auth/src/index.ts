import Fastify from 'fastify'

const app = Fastify({ logger: true })
const PORT = Number(process.env.AUTH_PORT) || 3001

app.get('/healthz', async () => ({ status: 'ok' }))
app.get('/readyz', async () => ({ status: 'ready' }))

// Routes (заглушки)
// POST /v1/auth/register
// POST /v1/auth/login
// POST /v1/auth/refresh
// POST /v1/auth/logout
// GET  /v1/auth/me

app.listen({ port: PORT, host: '0.0.0.0' }, (err) => {
  if (err) { app.log.error(err); process.exit(1) }
})
