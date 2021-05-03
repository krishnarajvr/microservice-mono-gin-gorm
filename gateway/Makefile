.PHONY: build run

AWS_ACCESS_KEY_ID=dummy-key
AWS_SECRET_ACCESS_KEY=dummy-secret
PERMISSION_URL=http://localhost:8082/api/v1/authorize
ACCOUNT_SERVICE=http://localhost:8082
PRODUCT_SERVICE=http://localhost:8083

export AWS_ACCESS_KEY_ID
export AWS_SECRET_ACCESS_KEY
export PERMISSION_URL
export ACCOUNT_SERVICE
export PRODUCT_SERVICE

build:
	 FC_ENABLE=1 FC_SETTINGS="$(PWD)/settings" \
	 FC_PARTIALS="$(PWD)/partials" \
	 FC_TEMPLATES="$(PWD)/templates" \
	 FC_OUT=$(PWD)/krakend-out.json \
	 krakend check -c $(PWD)/krakend.json -d

run:
	sed -i 's,http://account-service:8082,$(ACCOUNT_SERVICE),g' $(PWD)/krakend-out.json	
	sed -i 's,http://product-service:8083,$(PRODUCT_SERVICE),g' $(PWD)/krakend-out.json	

	krakend run -c $(PWD)/krakend-out.json -d

build-docker:
	sudo docker run --rm -it -v $(PWD):/etc/krakend -e FC_ENABLE=1 -e FC_SETTINGS="/etc/krakend/settings" -e FC_PARTIALS="/etc/krakend/partials" -e FC_OUT=out.json devopsfaith/krakend run -c /etc/krakend/krakend.json -d

run-docker:
	sudo docker run -p 8080:8080 -v $(PWD):/etc/krakend/ devopsfaith/krakend run -c /etc/krakend/krakend-out.json 	

replace:
	sed -i 's,http://account-service:8082,$(ACCOUNT_SERVICE),g' $(PWD)/krakend-out.json	

	

