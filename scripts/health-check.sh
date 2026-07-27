#!/bin/bash

set -e

kubectl get pods
kubectl get svc

curl http://localhost:8080/health || true
