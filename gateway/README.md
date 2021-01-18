# Gateway integration using Krakend

## Steps

### 1. Add service configuration to krakend.json 

### 2. Start the service using the configuration

```
docker run -p 8080:8080 -v $PWD:/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/krakend.json 
```


