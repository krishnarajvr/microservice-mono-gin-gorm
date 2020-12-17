#!/bin/sh
echo 'Runing migrations...'
/micro/bin/migrate up > /dev/null 2>&1 &

echo 'Start application...'
/micro/bin/api