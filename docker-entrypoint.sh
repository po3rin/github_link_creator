#!/usr/bin/env bash

# export GITHUB_CLIRNT_ID=$(aws ssm get-parameters --name ${PARAMETER_STORE_PREFIX}/clientid --query "Parameters[0].Value" --region ap-northeast-1 --output text)
# export GITHUB_SECRETGITHUB_SECRET=$(aws ssm get-parameters --name ${PARAMETER_STORE_PREFIX}/clientsecret --with-decryption --query "Parameters[0].Value" --region ap-northeast-1 --output text)
# env

set -e

export GITHUB_CLIENT_ID=$(aws ssm get-parameters --name /ghlinkcard/github/clientid --query "Parameters[0].Value" --region ap-northeast-1 --output text)
export GITHUB_SECRET=$(aws ssm get-parameters --name /ghlinkcard/github/clientsecret --with-decryption --query "Parameters[0].Value" --region ap-northeast-1 --output text)

exec "$@"