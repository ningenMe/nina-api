# nina-api

- how to build proto

```shell
cd proto
buf lint
npx buf generate nina
```

- how to set env-variable

```shell
`aws ssm get-parameters-by-path --path "/" --region ap-northeast-1 --output text | awk '{print "export",$5"="$7}'`
```

- curl

```shell
curl -XPOST -H 'Content-Type: application/json' -d '{}' localhost:8081/nina.v1.HealthService/Check -i
curl -XPOST -H 'Content-Type: application/json' -d '{}' https://nina-api.ningenme.net/nina.v1.HealthService/Check -i 
```
