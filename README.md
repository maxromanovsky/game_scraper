# Game Scraper

Warning: this is a pet project :)
Purpose is to play around with go, while checking what games that I've purchased on certain marketplaces (e.g. gog.com) disappeared from my account.

## Todo
- Multiple binaries
    - https://github.com/prometheus/prometheus/blob/master/cmd/promtool/main.go
    - https://github.com/golang-standards/project-layout
- Parquet or smth similar
- CLI arg parser
- Logger
- Test on ARM
- Defer to close resources
- Constructor?
- Struct inheritance patterns
- Testing
- Bazel build
- Upgrade dependencies:
    - https://golang.cafe/blog/upgrade-dependencies-golang.html
    - https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies
- Concurrency
    -https://medium.com/@matryer/golang-advent-calendar-day-two-starting-and-stopping-things-with-a-signal-channel-f5048161018
    - https://www.programming-books.io/essential/go/wait-for-goroutines-to-finish-ea3629ac73bb494283d0c92b2a4f78d1
    - https://blog.golang.org/pipelines


## Dep update
```shell script
go get -u ./... # all!!!
go get golang.org/x/net
go get golang.org/x/oauth2
go get google.golang.org/api
```