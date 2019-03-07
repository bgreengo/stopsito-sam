#!/usr/bin/env bash

GOOS=linux go build ./lambda-functions/src/create-reservation/main.go
zip create-reservation.zip main
rm main

GOOS=linux go build ./lambda-functions/src/publish-reservation-event/main.go
zip publish-reservation-event.zip main
rm main

aws cloudformation deploy --template-file cloudformation/create-s3-buckets.template --profile stopsito --stack-name stopsito-s3-buckets
aws s3 cp create-reservation.zip s3://sto-psito-lambda-functions/ --profile stopsito
aws s3 cp publish-reservation-event.zip s3://sto-psito-lambda-functions/ --profile stopsito

rm create-reservation.zip
rm publish-reservation-event.zip