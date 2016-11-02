# covman

## Benchmarks

```
This is ApacheBench, Version 2.3 <$Revision: 1748469 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /organizations
Document Length:        125518 bytes

Concurrency Level:      1
Time taken for tests:   17.913 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      125614000 bytes
HTML transferred:       125518000 bytes
Requests per second:    55.83 [#/sec] (mean)
Time per request:       17.913 [ms] (mean)
Time per request:       17.913 [ms] (mean, across all concurrent requests)
Transfer rate:          6848.20 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:    16   18   1.2     17      33
Waiting:       16   18   1.2     17      33
Total:         16   18   1.2     18      34

Percentage of the requests served within a certain time (ms)
  50%     18
  66%     18
  75%     18
  80%     18
  90%     19
  95%     20
  98%     21
  99%     22
 100%     34 (longest request)
```
