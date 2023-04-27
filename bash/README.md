# Purpose

This script leverages awscli and massages the JMESPath output to write parameters out to a file `.env`

## Requirements

Awscli is set up and you have sufficient IAM permissions. i.e. `ssm:GetParametersByPath`

## Usage

`SSM_PATH=/path/to/parameters/ ./paramstore.sh`

## Limitations

As there is plenty of string manipulation in the script, it is usually more accurate to use other languages leveraging AWS SDK