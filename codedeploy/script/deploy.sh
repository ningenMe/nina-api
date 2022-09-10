#!/bin/bash
aws ecr get-login-password --region AWS_REGION | docker login --username AWS --password-stdin AWS_ECR_REGISTRY
docker pull AWS_ECR_REGISTRY/AWS_ECR_REPOSITORY:latest
docker container stop nina-api | true
docker container rm nina-api | true
docker run -d -p8081:8081 --name nina-api AWS_ECR_REGISTRY/AWS_ECR_REPOSITORY:latest
