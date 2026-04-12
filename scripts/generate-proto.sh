#!/usr/bin/env bash
set -euo pipefail

echo "=== Generating proto files ==="

# Требует: protoc, protoc-gen-go, protoc-gen-go-grpc
for proto_dir in services/*/api/proto; do
  if ls "$proto_dir"/*.proto 2>/dev/null; then
    echo "Generating from $proto_dir..."
    protoc \
      --go_out=. \
      --go_opt=paths=source_relative \
      --go-grpc_out=. \
      --go-grpc_opt=paths=source_relative \
      "$proto_dir"/*.proto
  fi
done

echo "Done"
