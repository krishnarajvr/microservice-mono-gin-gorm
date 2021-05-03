FROM golang:1.16

WORKDIR /micro

RUN apt-get update && \
	apt-get install -y ca-certificates && \
	update-ca-certificates && \
    apt-get install git -y && \
	rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/devopsfaith/krakend-ce.git

WORKDIR /micro/krakend-ce

RUN go mod vendor

RUN make build

WORKDIR /micro

COPY . .

WORKDIR /micro/plugins

RUN go build -buildmode=plugin -o router-plugin.so ./router-plugin/*.go && \
   chmod +x /micro/deploy/bin/*

WORKDIR /micro

EXPOSE 8080



