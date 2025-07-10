#!/bin/bash

echo "[INFO] 🐳 Installing Docker CLI..."
apt-get update && apt-get install -y docker.io

echo "[INFO] 🟢 Installing Node.js..."
curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
apt-get install -y nodejs

echo "[INFO] ✅ Docker version:"
docker --version

echo "[INFO] ✅ Node.js version:"
node --version

echo "[INFO] ✅ npm version:"
npm --version

echo "[INFO] 🚀 Starting Jenkins..."
# Hand control back to the original Jenkins entrypoint
exec /usr/bin/tini -- /usr/local/bin/jenkins.sh
# exec /usr/local/bin/jenkins.sh
