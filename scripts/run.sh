#!/bin/bash
set -m

cleanup() {
    echo "Stopping all background processes..."
    kill $(jobs -p) 2>/dev/null
    wait
}

trap cleanup SIGINT SIGTERM

bash scripts/tunnel.sh &
bash scripts/run_go.sh &
bash scripts/run_web.sh &

wait
