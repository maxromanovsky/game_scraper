# Game Scraper

Warning: this is a pet project :)
Purpose is to play around with go, while checking what games that I've purchased on certain marketplaces (e.g. gog.com) disappeared from my account.

```
Items on your wishlist are now discounted!
Produkty z Twojej listy życzeń zostały przecenione!
A game you’ve been waiting for is now available on GOG.com
A gift for you!
Otrzymujesz prezent!
Hey, your free game has arrived!

Your order # xx is complete!
Twoje zamówienie nr xx zostało zrealizowane!
Free items added to your GOG.com library.
Do Twojej biblioteki GOG.com dodano darmowe produkty.
https://www.gog.com/email/preview/[a-z0-9]
```

## Todo
- CLI arg parser
- Logger with log levels
- Test on ARM
- Testing
- Bazel build
- Upgrade dependencies:
    - https://golang.cafe/blog/upgrade-dependencies-golang.html
    - https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies
- Concurrency
    -https://medium.com/@matryer/golang-advent-calendar-day-two-starting-and-stopping-things-with-a-signal-channel-f5048161018
    - https://www.programming-books.io/essential/go/wait-for-goroutines-to-finish-ea3629ac73bb494283d0c92b2a4f78d1
    - https://blog.golang.org/pipelines
- Multiple binaries
    - https://github.com/prometheus/prometheus/blob/master/cmd/promtool/main.go
- https://github.com/golang-standards/project-layout

## Dep update
```shell script
go get -u ./... # all!!!
go get golang.org/x/net
go get golang.org/x/oauth2
go get google.golang.org/api
go get github.com/actgardner/gogen-avro/v7
```

## Avro

- https://github.com/actgardner/gogen-avro#installation
- https://avro.apache.org/docs/current/spec.html

### Install CLI tool
Run outside of project dir, otherwise `go.mod` will be affected
```shell script
go get github.com/actgardner/gogen-avro/v7/cmd/...
go install github.com/actgardner/gogen-avro/v7/cmd/...
```

## Build

```shell script
go generate github.com/maxromanovsky/game_scraper/domain/entity
go build -o build github.com/maxromanovsky/game_scraper/cmd/mail_scraper
go build -o build github.com/maxromanovsky/game_scraper/cmd/mail_parser
```