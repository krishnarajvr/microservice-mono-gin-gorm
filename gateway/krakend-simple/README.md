# Gateway integration using Krakend

## Steps

### 1. Add service configuration to krakend.json ( Replace host with the service ip or domain)

### 2. Start the service using the configuration

```
make run 

or 

sudo docker run -p 8080:8080 -v $PWD:/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/krakend.json 
```


