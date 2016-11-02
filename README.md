# covman

## Benchmarks

```
This is ApacheBench, Version 2.3 <$Revision: 1748469 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/
          
Benchmarking localhost (be patient)
Completed 200 requests
Completed 400 requests
Completed 600 requests
Completed 800 requests
Completed 1000 requests
Completed 1200 requests
Completed 1400 requests
Completed 1600 requests
Completed 1800 requests
Completed 2000 requests
Finished 2000 requests
          
         
Server Software:
Server Hostname:        localhost
Server Port:            8080
          
Document Path:          /organizations
Document Length:        12 bytes
          
Concurrency Level:      100
Time taken for tests:   14.522 seconds
Complete requests:      2000
Failed requests:        1737
   (Connect: 0, Receive: 0, Length: 1737, Exceptions: 0)
Total transferred:      218225182 bytes
HTML transferred:       218027922 bytes
Requests per second:    137.72 [#/sec] (mean)
Time per request:       726.085 [ms] (mean)
Time per request:       7.261 [ms] (mean, across all concurrent requests)
Transfer rate:          14675.31 [Kbytes/sec] received
          
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1  36.4      0    1153
Processing:     6  717 391.2    705    2358
Waiting:        6  716 391.2    705    2358
Total:          9  718 393.3    706    2358

Percentage of the requests served within a certain time (ms)
  50%    706
  66%    865
  75%    957
  80%   1023
  90%   1228
  95%   1385
  98%   1602
  99%   1753
 100%   2358 (longest request)
```
