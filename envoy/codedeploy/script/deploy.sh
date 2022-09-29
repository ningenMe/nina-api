#!/bin/bash
aws ecr get-login-password --region AWS_REGION | docker login --username AWS --password-stdin AWS_ECR_REGISTRY

docker pull AWS_ECR_REGISTRY/AWS_ECR_REPOSITORY:latest
docker container stop nina-envoy-api | true
docker container rm nina-envoy-api | true
docker run -d -p8082:8082 --name nina-envoy-api --add-host host.docker.internal:host-gateway AWS_ECR_REGISTRY/AWS_ECR_REPOSITORY:latest
docker rmi `docker image ls | grep none | awk '{print $3}'` | true
