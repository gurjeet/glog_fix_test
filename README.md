Steps to Reproduce
==================

```
$ git clone --recursive https://github.com/gurjeet/glog_fix_test

$ go build -o main && ./main
./main flag redefined: v
panic: ./main flag redefined: v
... stack trace follows
```

