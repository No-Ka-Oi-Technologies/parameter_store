# Purpose

This script will loop through SSM Parameter Store with user-input path/ and write out to a file `.env`

## Requirements

The compiled binary only requires that your awscli is set up and you have sufficient IAM permissions. i.e. `ssm:GetParametersByPath`

## Usage

`./paramstore /path/to/parameters`