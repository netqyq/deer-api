```
shell$ ab -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ldHF5cUAxNjMuY29tIiwibmJmIjoxNDQ0NDc4NDAwfQ.bZo1DzrzZBetB9IP7fVip5XA_GiFBb_z8zDNTalReuU" -n 10000 -c 200 "localhost:9001/products/1"
This is ApacheBench, Version 2.3 <$Revision: 1663405 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9001

Document Path:          /products/1
Document Length:        84 bytes

Concurrency Level:      200
Time taken for tests:   26.599 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      3460000 bytes
HTML transferred:       840000 bytes
Requests per second:    375.96 [#/sec] (mean)
Time per request:       531.975 [ms] (mean)
Time per request:       2.660 [ms] (mean, across all concurrent requests)
Transfer rate:          127.03 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.8      0       7
Processing:     1  529 266.9    502    1622
Waiting:        1  529 266.9    502    1622
Total:          1  529 267.0    503    1622

Percentage of the requests served within a certain time (ms)
  50%    503
  66%    642
  75%    722
  80%    766
  90%    883
  95%    981
  98%   1118
  99%   1209
 100%   1622 (longest request)
```