## Steps to Execution

> Disclaimer: All steps below worked on Windows 11 Pro with Powershell on version 5.1

1- Install Kind, Kubens and Kubectx

```
Kind: https://kind.sigs.k8s.io/docs/user/quick-start/#installation
Kubens and Kubectx:  https://github.com/ahmetb/kubectx
```

2- Download repo desafio:

```
git clone https://github.com/desafioaltbank/desafio.git
```

3- Create a cluster kind in ./argocd/cluster.yaml:

```
kind create cluster --name cluster-alt-bank --config=argocd/kind_cluster.yaml
```

4- Create namespace argocd:

```
kubectl create namespace argocd
```

5- Install ArgoCD

```
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

6- Install CLI ArgoCD:

```
Visit: https://argo-cd.readthedocs.io/en/stable/cli_installation/#download-with-powershell-invoke-webrequest
```

7- Create a port-forwad to access argocd in the broswer:

```
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

8- Retrieve password initial from argocd installation:

```
$passwordArgo=kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}"
[System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($passwordArgo))
```

9- Login Argocd in http://localhost:8080 with user: admin and password retrieved in the step 8

10- Login by terminal with:

```
argocd login localhost:8080
```

11- Retrieve context with:

```
kubectl config current-context
```

12- Pass context to argocd with:

```
argocd cluster add SEU_CONTEXTO --insecure --in-cluster -y --upsert add
```

13- Confirm if cluster kind was add:

```
argocd cluster list
```

14- Retrieve then address your cluster with:

```
argocd cluster list
```

15- Test your cluster with argoCD, use the repo demo nginx:

```
kubens default
argocd app create nginx-app --repo https://github.com/renatovieiradesouza/k8s-deploy-nginx-example.git --path . --sync-policy automated --sync-retry-limit 5 --self-heal --auto-prune --dest-server https://kubernetes.default.svc --dest-namespace default
```

16- Confirm your app is create:

```
argocd app list and argocd app get nginx-app
```

17- Test with port-forward

```
kubectl port-forward svc/nginx-svc 9090:9113
```

access: [http://localhost:9090/metrics](http://localhost:9090/metrics)

## Working with a manifest argocd

1- In repository contain manifest to configure repository git, build image, math app code and github actions code
2- Retrieve then address your cluster with:

```
argocd cluster list
```

3- Create a ns to apimath - Prime Number API

```
kubectl  create ns apimath
```

4- Create a app in argocd:

```
kubens argocd

argocd app create api-math --repo https://github.com/desafioaltbank/desafio.git --path ./k8s --dest-server https://kubernetes.default.svc --sync-policy automated --sync-retry-limit 5 --self-heal --auto-prune --dest-namespace apimath
```

5- Now this repo be linked with argocd, change your app, make a push and test automated sync

6- Test api with:

```
kubectl port-forward svc/apimath -n apimath 8282:8080

Powershell
$headers = New-Object "System.Collections.Generic.Dictionary[[String],[String]]"
$headers.Add("Content-Type", "application/json")

$body = @"
{`"number`": 104743}
"@

$response = Invoke-RestMethod 'http://localhost:8282/prime' -Method 'POST' -Headers $headers -Body $body
$response | ConvertTo-Json

Linux

curl -X POST -H "Content-Type: application/json" -d '{"number": 104743}' http://localhost:8282/prime

```

7- Test the pipeline in Github Actions send a PR to execute the action:

access: [github.com/desafioaltbank/desafio/actions](https://github.com/desafioaltbank/desafio/actions)

7- Destroy cluster:

```
kind delete clusters cluster-alt-bank
```
