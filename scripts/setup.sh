#!/usr/bin/env bash
set -euo pipefail

echo "=== AutoAPI-AK-Core: Setup ==="

# Копируем .env файлы
cp -n .env.example .env 2>/dev/null && echo "✅ .env created" || echo "⚠️  .env already exists"

for svc in services/*/; do
  if [ -f "$svc.env.example" ]; then
    cp -n "$svc.env.example" "$svc.env" 2>/dev/null && echo "✅ $svc.env created" || echo "⚠️  $svc.env already exists"
  fi
done

echo ""
echo "Next steps:"
echo "  1. Fill in all .env files with your values"
echo "  2. Run: make up"
