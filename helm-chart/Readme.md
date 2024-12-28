## Helm it up

```
helm plugin install https://github.com/databus23/helm-diff


helm create gman-go-server

kubectl create ns gman-go-server

helm install --dry-run --debug --namespace gman-go-server sgune-server ./gman-go-server  # dry run

helm upgrade --install --cleanup-on-fail --namespace gman-go-server sgune-server -f gman-go-server/values/gke.yaml  ./gman-go-server

helm diff upgrade --namespace gman-go-server sgune-server -f gman-go-server/values/gke.yaml  ./gman-go-server -C3

#NOTES
Get the application URL by running these commands: 
  export POD_NAME=$(kubectl get pods --namespace gman-go-server -l "app.kubernetes.io/name=gman-go-server,app.kubernetes.io/instance=sgune" -o jsonpath="{.items[0].metadata.name}")
  export CONTAINER_PORT=$(kubectl get pod --namespace gman-go-server $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace gman-go-server port-forward $POD_NAME 8080:$CONTAINER_PORT

```