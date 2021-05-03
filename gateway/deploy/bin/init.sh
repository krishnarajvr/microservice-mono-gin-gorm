#!/bin/sh
echo 'Runing Gateway...'

sed -i "s,http://account-service:8082,$ACCOUNT_SERVICE,g" /micro/krakend.json
sed -i "s,http://content-service:8083,$PRODUCT_SERVICE,g" /micro/krakend.json
sed -i "s,http://asset-service:8084,$ASSET_SERVICE,g" /micro/krakend.json
sed -i "s,http://vendor-service:8085,$VENDOR_SERVICE,g" /micro/krakend.json
sed -i "s,http://aggregate-service:8086,$AGGREGATE_SERVICE,g" /micro/krakend.json
/micro/krakend run -c /micro/krakend.json
