# istio-example
go/golang istio example

```bash
# 创建`example`命名空间
kubectl create namespace example
# 删除`example`命名空间
kubectl delete namespace example
# 获取所有命名空间
kubectl get namespace

# 查看`example`命名空间中的Pod
kubectl get pod -n example -owide

# 给`example`命名空间设置自动 Sidecar 注入
kubectl label namespace example istio-injection=enabled

# 获取 kubernetes-dashboard token
kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')

# Istio
```bash
kubectl get svc -n istio-system
kubectl get pods -n istio-system
```

# 在example命名空间中安装server
kubectl apply -n example -f https://raw.githubusercontent.com/DeyiXu/istio-example/master/server/server.yaml
# 在example命名空间中删除server
kubectl delete -n example -f https://raw.githubusercontent.com/DeyiXu/istio-example/master/server/server.yaml

# 在example命名空间中安装client
kubectl apply -n example -f https://raw.githubusercontent.com/DeyiXu/istio-example/master/client/client.yaml
# 在example命名空间中删除client
kubectl delete -n example -f https://raw.githubusercontent.com/DeyiXu/istio-example/master/client/client.yaml

# 在example命名空间中安装client-gateway
kubectl apply -n example -f https://raw.githubusercontent.com/DeyiXu/istio-example/master/client/client-gateway.yaml
# 在example命名空间中删除client-gateway
kubectl delete -n example -f https://raw.githubusercontent.com/DeyiXu/istio-example/master/client/client-gateway.yaml
```