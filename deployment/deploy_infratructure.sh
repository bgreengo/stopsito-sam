#!/usr/bin/env bash

AWS_PROFILE=$1

echo "[+] Creating DynamoDb Tables"
{
    aws cloudformation deploy --template-file ./cloudformation/create-dynamodb-tables.template --profile ${AWS_PROFILE} --stack-name stopsito-dynamodb-tables
} &> /dev/null

echo "[+] Creating Api Gateway Endpoints"
{
    aws cloudformation deploy --template-file ./cloudformation/create-api-gateway-endpoints.template --profile ${AWS_PROFILE} --stack-name stopsito-api-endpoints
} &> /dev/null

echo "[+] Creating SNS Topics"
{
    aws cloudformation deploy --template-file ./cloudformation/create-sns-topics.template --profile ${AWS_PROFILE} --stack-name stopsito-sns-topics
} &> /dev/null

echo "[!] Finished!"