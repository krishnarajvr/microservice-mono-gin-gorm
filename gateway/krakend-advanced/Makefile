.PHONY: run

build:
	sudo docker run --rm -it -v $(PWD):/etc/krakend -e FC_ENABLE=1 -e FC_SETTINGS="/etc/krakend/settings" -e FC_PARTIALS="/etc/krakend/partials" -e FC_OUT=out.json devopsfaith/krakend run -c /etc/krakend/krakend.json -d

run:
	sudo docker run -p 8080:8080 -v $(PWD):/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/out.json 

