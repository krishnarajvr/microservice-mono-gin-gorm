# Gateway integration using Krakend

## Requirements

* Golang - 1.15 recommended
* [krakend](https://www.krakend.io/)  

## Build krakend locally

1. Download source code from https://github.com/devopsfaith/krakend-ce/
2. Go to source folder and run ```make build```
3. Copy the ```krakend``` binary to go bin folder ```{GOPATH}/bin```

### Steps

Build Kradend binary locally


```
git clone https://github.com/devopsfaith/krakend-ce.git

cd krakend-ce

git checkout v1.3.0

go mod vendor

make build

```

Copy the kradend binary to  ```{GOPATH}/bin```  path once created.

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
Make sure current directory is ./gateway . Not the plugin directory. ( ``` cd ../../ ``` )

It will generate krakend-out.json file with final configuration from the templates and settings. 

### Start the gateway

#### Change the configurations in Makefile as per the setup of other services
```
PERMISSION_URL=http://localhost:8082/api/v1/authorize
ACCOUNT_SERVICE=http://localhost:8082
PRODUCT_SERVICE=http://localhost:8083
```

#### Run the gateway

```
make run
```

If there is any configuration change in krakend.json, It need to build the chagnes again using ``` make build ``` in gateway folder. Refer Makefile for the commands that is running.

Note : check ```{GOPATH}/bin``` should contain krakend binary as per the krakend-ce build
