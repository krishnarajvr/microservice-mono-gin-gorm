# Build environment
# -----------------
FROM golang:1.15-buster as build-env
WORKDIR /micro

RUN apt-get update && \
    apt-get install git -y 

RUN git clone https://github.com/devopsfaith/krakend-ce.git

WORKDIR /micro/krakend-ce

RUN go mod vendor

RUN make build

WORKDIR /micro

RUN pwd

COPY . .


RUN  FC_ENABLE=1 FC_SETTINGS=settings FC_PARTIALS=partials FC_TEMPLATES=templates FC_OUT=krakend-out.json /micro/krakend-ce/krakend check -c krakend.json -d

WORKDIR /micro/plugins/router-plugin

RUN go build -buildmode=plugin -o ../router-plugin.so ./*.go

WORKDIR /micro

# Deployment environment
# ----------------------
FROM debian:bullseye-slim
RUN apt-get update && \
    apt-get install ca-certificates -y && \
    update-ca-certificates

WORKDIR /micro

COPY --from=build-env /micro/krakend-ce/krakend /micro/
COPY --from=build-env /micro/plugins/*.so /micro/plugins/
COPY --from=build-env /micro/krakend-out.json /micro/krakend.json
COPY --from=build-env /micro/deploy/bin/init.sh /micro/bin/init.sh

RUN chmod +x /micro/bin/*

EXPOSE 8080 8090
