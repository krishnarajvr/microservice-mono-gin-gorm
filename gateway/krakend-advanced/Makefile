.PHONY: run build

build:
	 FC_ENABLE=1 FC_SETTINGS="$(PWD)/settings" FC_PARTIALS="$(PWD)/partials" FC_OUT=$(PWD)/out.json ~/work/aha/krakend-ce/krakend run -c $(PWD)/krakend.json -d

run:
	~/work/aha/krakend-ce/krakend run -c $(PWD)/out.json -d

build-docker:
	sudo docker run --rm -it -v $(PWD):/etc/krakend -e FC_ENABLE=1 -e FC_SETTINGS="/etc/krakend/settings" -e FC_PARTIALS="/etc/krakend/partials" -e FC_OUT=out.json devopsfaith/krakend run -c /etc/krakend/krakend.json -d

run-docker:
	sudo docker run -p 8080:8080 -v $(PWD):/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/out.json 	

