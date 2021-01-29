.PHONY: run

run-docker:
	sudo docker run -p 8080:8080 -v $(PWD):/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/krakend.json 
run: 
	~/work/aha/krakend-ce/krakend run -c krakend.json -d
plugin:
	go build -buildmode=plugin -o github-data.so ./plugin
plugin-example:
	go build -buildmode=plugin -o client-example.so ./plugin	