#!/usr/bin/env bash
kubectl create -f ./deployments/mongodb-ssd.yml
kubectl create -f ./deployments/mongodb-deployment.yml
kubectl create -f ./deployments/mongodb-service.yml
