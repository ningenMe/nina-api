# nina-api

## go get
```shell
go get -u github.com/ningenMe/mami-interface@v1.x.0
```

```shell
`aws ssm get-parameters-by-path --path "/" --region ap-northeast-1 --output text | awk '{print "export",$5"="$7}'`

```

## endpoint
```shell
grpcurl -plaintext localhost:8081 list
grpcurl nina-api.ningenme.net:443 list 

grpcurl -plaintext localhost:8081 nina.HealthService/Get
grpcurl nina-api.ningenme.net:443 nina.HealthService/Get
```

### github contribution

```shell
grpcurl -plaintext localhost:8081 nina.GithubContributionService.Get
grpcurl nina-api.ningenme.net:443 nina.GithubContributionService.Get
```

```shell
grpcurl -plaintext -d '
    {
      "contribution" : {
        "contributedAt": "2022-09-14T00:00:00+09:00",
        "organization": "org1",
        "repository": "repo1",
        "user": "user1",
        "status": "status1"
      }
    }
    {
      "contribution" : {
        "contributedAt": "2022-09-14T00:00:00+09:00",
        "organization": "org2",
        "repository": "repo2",
        "user": "user2",
        "status": "status2"
      }
    }
' localhost:8081 nina.GithubContributionService.Post
```

```shell
grpcurl -plaintext -d '
    {
      "startAt" : "2022-09-14T00:00:00+09:00",
      "endAt"   : "2022-09-16T00:00:00+09:00"
    }
' localhost:8081 nina.GithubContributionService.Delete
grpcurl -d '
    {
      "startAt" : "2022-09-14T00:00:00+09:00",
      "endAt"   : "2022-09-16T00:00:00+09:00"
    }
' nina-api.ningenme.net:443 nina.GithubContributionService.Delete
```

```shell
grpcurl -plaintext -d '
    {
      "user" : "ningenMe"
    }
' localhost:8081 nina.GithubContributionService.GetStatistics
grpcurl -d '
    {
      "user" : "ningenMe"
    }
' nina-api.ningenme.net:443 nina.GithubContributionService.GetStatistics
```

### compro category

```shell
grpcurl -plaintext localhost:8081 nina.ComproCategoryService/Get
grpcurl nina-api.ningenme.net:443 nina.ComproCategoryService/Get
```

```shell
grpcurl -plaintext -d '
    {
      "category" : {
        "categoryDisplayName": "テスト",
        "categorySystemName": "test",
        "categoryOrder": 1
      }
    }
' localhost:8081 nina.ComproCategoryService/Post
grpcurl -plaintext -d '
    {
      "categoryId" : "category_6H8BTC",
      "category" : {
        "categoryDisplayName": "テスト改",
        "categorySystemName": "test2",
        "categoryOrder": -1
      }
    }
' localhost:8081 nina.ComproCategoryService/Post
grpcurl -plaintext -d '
    {
      "categoryId" : "category_6H8BTC"
    }
' localhost:8081 nina.ComproCategoryService/Post
```