#!/bin/bash

set -e

kubectl rollout undo deployment/backend-api
kubectl rollout status deployment/backend-api
