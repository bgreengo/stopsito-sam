#!/usr/bin/env bash
aws cloudformation deploy --template-file cloudformation/setup.template --profile stopsito --stack-name stopsito-sam