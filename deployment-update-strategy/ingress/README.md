# Canary Releases in 15 Minutes

> Demo for Docker for Mac with Kuberentes

## Setting up Nginx Ingress on Kubernetes on Docker for Mac with Kuberentes

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml
```

## Clone repo

```
git clone https://github.com/kaluzaaa/k8s-canary-on-ingress.git
cd k8s-canary-on-ingress
```

## Deploy app

```
kubectl apply -f web/deployment.yaml
kubectl apply -f cat/deployment.yaml
kubectl apply -f dog/deployment.yaml
```

## Deploy ingress rule for web app and backend

```
kubectl apply -f ingress/1-ingress.yaml
```

Check if the api works.

```
curl localhost/pet
```

Check web app in browser at [http://localhost/web](http://localhost/web)

## Deploy canary ingress rule and test

Split traffic 90:10.

```
kubectl apply -f ingress/2-canary-10.yaml
```

Check with curl or web app.

```
curl localhost/pet
```

Check 100 request with ruby script.

```
ruby ingress/check.rb
```

Set traffic 50:50.

```
kubectl apply -f ingress/3-canary-50.yaml
```

Check 100 request with ruby script.

```
ruby ingress/check.rb
```

Remove rule.

```
kubectl delete -f ingress/3-canary-50.yaml
```

## A/B test with header

Deploy canary ingress rule.

```
kubectl apply -f ingress/4-canary-header.yaml
```

Check if the api works.

```
curl localhost/pet
```

```
curl  -H "dog-lover: always" localhost/pet
```

Check web app in browser at [http://localhost/web](http://localhost/web)

Remove rule.

```
kubectl delete -f ingress/4-canary-header.yaml
```

## A/B test with cookie

Deploy canary ingress rule.

```
kubectl apply -f ingress/5-canary-cookie.yaml
```

Check if the api works.

```
curl localhost/pet
```

```
curl  -b dog_lover=always http://localhost/pet
```

Check web app in browser at [http://localhost/web](http://localhost/web)

Remove rule.

```
kubectl delete -f ingress/5-canary-cookie.yaml
```
