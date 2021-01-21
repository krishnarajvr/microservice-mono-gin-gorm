# Gateway integration using Krakend

## Steps

### 1. Add service configuration to trafiek.yaml ( Replace host with the service ip or domain)

### 2. Start the service using the configuration

```
make run 

or 

sudo docker run -p 8080:8080 -p 81:80  -v $PWD:/etc/traefik/ traefik:v2.4
```


