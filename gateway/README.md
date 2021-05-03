# Gateway integration using Krakend

## Requirements

* Golang - 1.15 recommended
* [krakend](https://www.krakend.io/)  

## Build krakend locally

1. Download source code from https://github.com/devopsfaith/krakend-ce/
2. Go to source folder and run ```make build```
3. Copy the ```krakend``` binary to go bin folder ```{GOPATH}/bin```

### Steps

Install Kradend binary locally


```
git clone https://github.com/devopsfaith/krakend-ce.git

cd krakend-ce

git checkout v1.3.0

go mod vendor

make build

```

## Build Plugins

Build router plugin 

```
cd plugins/router-plugin

make build

```

## dynamic krakend config

Refer : https://www.krakend.io/docs/configuration/flexible-config/


### Generate Kradend Config from dynamic configuration

```
make build
```

### Start the gateway

#### Change the configurations in Makefile as per the setup
```
PERMISSION_URL=http://localhost:8082/api/v1/authorize
ACCOUNT_SERVICE=http://localhost:8082
PRODUCT_SERVICE=http://localhost:8083
```

#### Run the gateway

```
make run
```

Note : check ```{GOPATH}/bin``` contain krakend binary