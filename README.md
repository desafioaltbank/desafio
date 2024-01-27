## Steps to Success

1- Install Kind

```
Visit: https://kind.sigs.k8s.io/docs/user/quick-start/#installation
```

2- Create a cluster kind in ./argocd/cluster.yaml:

```
kind create cluster --name cluster-alt-bank --config=argocd/kind_cluster.yaml
```

3- Create namespace argocd:

```
kubectl create namespace argocd
```

4- Install ArgoCD

```
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

4- Install CLI ArgoCD:

```
Visit: https://argo-cd.readthedocs.io/en/stable/cli_installation/#download-with-powershell-invoke-webrequest
```

5- Create a port-forwad to access argocd in the broswer:

```
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

6- Retrieve password initial from argocd installation:

```
kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d
```

7- Login Argocd in http://localhost:8080 with user: admin and password retrieved in the step 6

8- Login by terminal with:

```
argocd login localhost:8080
```

9- Retrieve context with:

```
kubectl config current-context
```

10- Pass context to argocd with:

```
argocd cluster add SEU_CONTEXTO --insecure --in-cluster -y --upsert add
```

11- Confirm if cluster kind was add:

```
argocd cluster list
```

12- Retrieve then address your cluster with:

```
argocd cluster list
```

13- Test your cluster with argoCD, use the repo demo nginx:

```
argocd app create nginx-app --repo https://github.com/renatovieiradesouza/k8s-deploy-nginx-example.git --path . --dest-server https://kubernetes.default.svc --dest-namespace default
```

14- Confirm your app is create:

```
argocd app list and argocd app get nginx-app
```

15- Sync your application in argoCD:

```
argocd app sync nginx-app
```

16- Enable auto sync to repo:

```
argocd app set nginx-app --sync-policy automated
```

17- Test with change in your repo, change number of replicas

## Working with a manifest argocd

1- In repository contain manifest to configure repository git, build image, math app code and github actions code
2- Retrieve then address your cluster with:

```
argocd cluster list
```

3- Create a ns to apimath

```
kubectl  create ns apimath
```

4- Create a app in argocd:

```
argocd app create api-math --repo https://github.com/desafioaltbank/desafio.git --path ./k8s --dest-server https://kubernetes.default.svc --dest-namespace apimath
```

5- Enable auto sync to repo:

```
argocd app set api-math --sync-policy automated
```

6- Now this repo be linked with argocd, change your app, make a push and test automated sync
