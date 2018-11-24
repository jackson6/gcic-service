#!/usr/bin/env bash
kubectl create secret generic cloudsql-instance-credentials --from-file=credentials.json=./private/gcic-219917-8578ad2ea5a6.json
kubectl create secret generic cloudsql-db-credentials --from-literal=username=postgres --from-literal=password=dreamer6