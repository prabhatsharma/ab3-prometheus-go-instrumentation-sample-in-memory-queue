#!/bin/sh

NAMESPACE=prometheus-sample
kubectl proxy --port 8009 &
kubectl get namespace $NAMESPACE -o json |jq '.spec = {"finalizers":[]}' >temp.json
curl -k -H "Content-Type: application/json" -X PUT --data-binary @temp.json 127.0.0.1:8009/api/v1/namespaces/$NAMESPACE/finalize
