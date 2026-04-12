import Fastify from 'fastify'

const app = Fastify({ logger: true })
const PORT = Number(process.env.GATEWAY_PORT) || 3000

app.get('/healthz', async () => ({ status: 'ok' }))
app.get('/readyz', async () => ({ status: 'ready' }))

// TODO: регистрировать плагины, маршруты, middleware

app.listen({ port: PORT, host: '0.0.0.0' }, (err) => {
  if (err) { app.log.error(err); process.exit(1) }
})
