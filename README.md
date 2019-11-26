# aws-serverless-golang-101

## Requirements

Install Node.js and AWS CLI in your PC.  
After that, execute `aws configure` and setup your access key to AWS in interactive command.

Install Serverless framework CLI in your PC
```bash
yarn global add serverless
sls -v
```

## How to deploy your Lambda functions to AWS

To deploy lambda functions to AWS
```bash
make deploy
```

To remove lambda functions from AWS
```bash
make remove
```

## Reference

[aws-golang-http-get-post](https://github.com/serverless/examples/tree/master/aws-golang-http-get-post)
