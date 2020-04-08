# taxi HTTP API

[![Build Status][circleci-badge]][circleci-link]
[![Report Card][report-badge]][report-link]

## Installation

```bash
$ go get github.com/Alma-media/taxi
```

## Docker
- Build:
```bash
$ docker build . -t taxi
```
- Run:
```bash
$ docker run --rm -it -p 127.0.0.1:8080:8080 taxi
```

## Furter improvements

### API
- enable the access control for public and private endpoints (use middleware)
- enable codec middleware and retrieve codec from the context to use corresponding encoder (e.g. `codec.NewEncoder(w).Encode(order)`)
- pass logger to the handler (intermal errors should be logged but not exposed to the user)
- find approximate limit to enable throttling

### Config
- find a better library to parse the values from flags, toml, yaml, json ...

### Repository
- switch to event driven approach and get rid of proxy (current repository implementation) to increase the throughput

### Storage
- choose another storage implementation that uses the context and is able to fail / return an error (e.g. database, blockchain etc)

## Benchmark results
```bash
[alma@alexsch taxi]$ ab -n 1000000 -c 1000 http://localhost:8080/request/

Server Softwa[![GoDoc][godoc-badge]][godoc-link]
re:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /request/
Document Length:        3 bytes

Concurrency Level:      1000
Time taken for tests:   37.403 seconds
Complete requests:      1000000
Failed requests:        0
Total transferred:      118975648 bytes
HTML transferred:       3000000 bytes
Requests per second:    26735.49 [#/sec] (mean)
Time per request:       37.403 [ms] (mean)
Time per request:       0.037 [ms] (mean, across all concurrent requests)
Transfer rate:          3106.32 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   18   5.4     17      45
Processing:     4   20   5.8     19      78
Waiting:        0   11   5.6      9      55
Total:         19   37   5.9     35     104

Percentage of the requests served within a certain time (ms)
  50%     35
  66%     36
  75%     38
  80%     39
  90%     47
  95%     51
  98%     55
  99%     58
 100%    104 (longest request)
```

[circleci-badge]: https://circleci.com/gh/Alma-media/taxi.svg?style=shield
[circleci-link]: https://circleci.com/gh/Alma-media/taxi
[report-badge]: https://goreportcard.com/badge/github.com/Alma-media/taxi
[report-link]: https://goreportcard.com/report/github.com/Alma-media/taxi
