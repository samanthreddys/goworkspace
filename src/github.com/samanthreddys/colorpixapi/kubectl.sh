#!/usr/bin/env bash
kubectl create -f db-service.yaml,db-deployment.yaml,colorpixapi-service.yaml,colorpixapi-deployment.yaml,colorpixapi-claim0-persistentvolumeclaim.yaml