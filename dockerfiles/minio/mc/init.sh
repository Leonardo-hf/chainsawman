#!/bin/bash
mc config host add minio http://minio-svc:9000 minioadmin minioadmin
mc mb --ignore-existing minio/source
mc ilm add --expiry-days "1" minio/source
mc mb --ignore-existing minio/algo
mc mb --ignore-existing minio/lib
for file in `ls /jar`
do
if test -f /jar/$file
then
  mc cp /jar/$file minio/lib
fi
done
