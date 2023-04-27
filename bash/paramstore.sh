#!/bin/bash -e

aws ssm get-parameters-by-path \
  --path $SSM_PATH \
  --with-decryption \
  --query "Parameters[*].[Name,Value]" \
  --output text | sed -e "s|$SSM_PATH||g" | sed -e 's/\t/=/g' > .env