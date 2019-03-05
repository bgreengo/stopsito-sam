#!/usr/bin/env bash

GOOS=linux go build ./lambda-functions/src/create-reservation/main.go
zip create-reservation.zip main
rm main

aws cloudformation deploy --template-file cloudformation/lambda-functions-bucket.template --profile stopsito --stack-name stopsito-lambda-bucket
aws s3 cp create-reservation.zip s3://sto-psito-lambda-functions/ --profile stopsito
rm create-reservation.zip