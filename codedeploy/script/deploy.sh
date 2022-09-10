#!/bin/bash
aws ecr get-login-password --region AWS_REGION | docker login --username AWS --password-stdin AWS_ECR_REGISTRY
docker pull AWS_ECR_REGISTRY/AWS_ECR_REPOSITORY:latest
