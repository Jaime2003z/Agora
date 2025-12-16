#!/bin/bash

# Detener cualquier nodo en ejecución
pkill -f agoranode

# Iniciar nodo 1 (bootstrap)
./agoranode --config=testnet/nodes/node1/config.toml --data-dir=testnet/nodes/node1/data --log-level=debug --name=node1 --port=3000 &

# Esperar un momento para que el primer nodo arranque
sleep 2

# Iniciar nodo 2
./agoranode --config=testnet/nodes/node2/config.toml --data-dir=testnet/nodes/node2/data --log-level=debug --name=node2 --port=3001 --bootstrap=localhost:3000 &

# Iniciar nodo 3
./agoranode --config=testnet/nodes/node3/config.toml --data-dir=testnet/nodes/node3/data --log-level=debug --name=node3 --port=3002 --bootstrap=localhost:3000 &

echo "Nodos iniciados. Usa 'pkill -f agoranode' para detenerlos."
