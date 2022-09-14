# nina-api

- go get
```shell
go get -u github.com/ningenMe/mami-interface@v0.x.0
```

- endpoint
```shell
grpcurl -plaintext localhost:8081 list
grpcurl -plaintext localhost:8081 nina.Health/Get
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
' localhost:8081 nina.GithubContribution.Post
```