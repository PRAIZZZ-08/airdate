#!/bin/bash
cd apps/api && go run ./cmd/server & \
cd apps/web && npm run dev
