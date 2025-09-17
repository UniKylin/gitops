# awesome-kubernetes

## Deployment

### 1. 怎样理解 `Deployment`
  管理`Pod` 管理很多 `Pod` 但两者之间又不是强绑定关系！

### 2. 如何查看 `Deployment` 的类型
```shell
$ kubectl api-resources
NAME                                SHORTNAMES   APIVERSION                        NAMESPACED   KIND
deployments                         deploy       apps/v1                           true         Deployment
```
### 3. 怎样生成描述文件
```shell
$ kubectl create deploy ngx-dep --image=nginx:alpine --dry-run=client -o yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: ngx-dep
  name: ngx-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ngx-dep
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: ngx-dep
    spec:
      containers:
      - image: nginx:alpine
        name: nginx
        resources: {}
status: {}
```

### 4. `Deployment` 的部署、查看、删除`Pod`、扩缩容、删除等操作
```shell
~/gitops/deploy on  master! ⌚ 11:37:25
$ kubectl apply -f dep.yaml
deployment.apps/ngx-dep created

~/gitops/deploy on  master! ⌚ 11:38:17
$ kubectl get deploy
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
ngx-dep   4/4     4            4           5s

~/gitops/deploy on  master! ⌚ 11:38:22
$ kubectl get pod
NAME                      READY   STATUS    RESTARTS   AGE
ngx-dep-77fcd74db-89g25   1/1     Running   0          10s
ngx-dep-77fcd74db-bghwl   1/1     Running   0          10s
ngx-dep-77fcd74db-jp8wm   1/1     Running   0          10s
ngx-dep-77fcd74db-q98hp   1/1     Running   0          10s

~/gitops/deploy on  master! ⌚ 11:38:27
$ kubectl delete pod ngx-dep-77fcd74db-89g25
pod "ngx-dep-77fcd74db-89g25" deleted

~/gitops/deploy on  master! ⌚ 11:39:20
$ kubectl get pod
NAME                      READY   STATUS    RESTARTS   AGE
ngx-dep-77fcd74db-bghwl   1/1     Running   0          66s
ngx-dep-77fcd74db-jp8wm   1/1     Running   0          66s
ngx-dep-77fcd74db-jvdhn   1/1     Running   0          4s
ngx-dep-77fcd74db-q98hp   1/1     Running   0          66s

~/gitops/deploy on  master! ⌚ 11:39:23
$ kubectl scale --replicas=2 deploy ngx-dep
deployment.apps/ngx-dep scaled

~/gitops/deploy on  master! ⌚ 11:40:48
$ kubectl get pod
NAME                      READY   STATUS    RESTARTS   AGE
ngx-dep-77fcd74db-bghwl   1/1     Running   0          2m33s
ngx-dep-77fcd74db-jvdhn   1/1     Running   0          91s

~/gitops/deploy on  master! ⌚ 11:40:50
$ kubectl get pod -o wide
NAME                      READY   STATUS    RESTARTS   AGE     IP           NODE              NOMINATED NODE   READINESS GATES
ngx-dep-77fcd74db-bghwl   1/1     Running   0          2m41s   10.244.1.5   desktop-worker2   <none>           <none>
ngx-dep-77fcd74db-jvdhn   1/1     Running   0          99s     10.244.3.6   desktop-worker3   <none>           <none>

~/gitops/deploy on  master! ⌚ 11:43:23
$ kubectl delete -f dep.yaml
deployment.apps "ngx-dep" deleted

$ kubectl get pod -o wide
No resources found in default namespace.
```

