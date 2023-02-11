#!/bin/bash

whoami
pwd
ls
cd /home/ec2-user/nina-api
kubectl apply -f deployment.yaml -n ningenme-space
kubectl apply -f service.yaml -n ningenme-space