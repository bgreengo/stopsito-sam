#!/usr/bin/env bash

AWS_PROFILE=$1

echo "[+] Building Lambda Functions"
{
    GOOS=linux go build ./lambda-functions/src/create-reservation/main.go
    zip create-reservation.zip main
    rm main

    GOOS=linux go build ./lambda-functions/src/publish-reservation-event/main.go
    zip publish-reservation-event.zip main
    rm main

    GOOS=linux go build ./lambda-functions/src/send-sms/main.go
    zip send-sms.zip main
    rm main

    GOOS=linux go build ./lambda-functions/src/send-confirmation-emails/main.go
    zip send-confirmation-emails.zip main
    rm main

} &> /dev/null

echo "[+] Creating S3 bucket"
{
    aws cloudformation deploy --template-file ./cloudformation/create-s3-buckets.template --profile ${AWS_PROFILE} --stack-name stopsito-s3-buckets
} &> /dev/null

echo "[+] Uploading Lambda Functions to S3"
{
    aws s3 cp create-reservation.zip s3://sto-psito-lambda-functions/ --profile ${AWS_PROFILE}
    aws s3 cp publish-reservation-event.zip s3://sto-psito-lambda-functions/ --profile ${AWS_PROFILE}
    aws s3 cp send-sms.zip s3://sto-psito-lambda-functions/ --profile ${AWS_PROFILE}
    aws s3 cp send-confirmation-emails.zip s3://sto-psito-lambda-functions/ --profile ${AWS_PROFILE}
} &> /dev/null

echo "[+] Cleaning up"
{
    rm create-reservation.zip
    rm publish-reservation-event.zip
    rm send-sms.zip
    rm send-confirmation-emails.zip
} &> /dev/null