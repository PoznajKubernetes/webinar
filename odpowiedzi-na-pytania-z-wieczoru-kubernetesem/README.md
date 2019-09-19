[Slajdy](slajdy.pdf)

## Demo

Wgranie przykładowej aplikacji - w tym wypadku nginx.

```
kubectl apply -f demo.yaml
```

Obserwowanie stanu podów

```
watch kubectl get pod -o wide
```

oraz węzłów klastra

```
watch kubectl get nodes
```

Symulacja awarii pod przez skasowanie jednego pod

```
kubectl delete pod nginx-XXX-YYY
```

Symulacja awarii noda poprzez jego "wyrwanie z prądu" na Azure

```
az vm stop -g <nazwa rg> -n <nazwa węzła> --skip-shutdown --no-wait
```

## Ingress

[Przykładowa reguła Ingress](ingress.yaml)