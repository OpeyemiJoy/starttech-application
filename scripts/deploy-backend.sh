#!/bin/bash

set -e

kubectl apply -f k8s/
kubectl rollout status deployment/backend-api
