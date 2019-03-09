#!/usr/bin/env bash

AWS_PROFILE=$1
./deployment/deploy_lambda_functions.sh ${AWS_PROFILE}
./deployment/deploy_infratructure.sh ${AWS_PROFILE}