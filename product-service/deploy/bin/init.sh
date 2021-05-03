#!/bin/sh
echo 'Runing migrations...'
cd micro;./migrate up > /dev/null 2>&1 &

echo 'Start application...'
/micro/api
