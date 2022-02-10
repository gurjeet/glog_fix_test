Steps to Reproduce
==================

```
$ git clone --recursive https://github.com/gurjeet/glog_fix_test

$ go build -o main && ./main
./main flag redefined: v
panic: ./main flag redefined: v
... stack trace follows
```

How to Test the Fix
===================

```
$ go mod edit -replace=github.com/golang/glog=./GLOG_ORIGIN

$ (cd GLOG_ORIGIN/ && patch -p1 < ../bugfix.patch)

$ go build -o main -ldflags "-X github.com/golang/glog.OptionPrefix=glog:"
$ ./main -glog:logtostderr -glog:v 1 -v somevalue

I0210 11:17:28.096799   72018 main.go:34] Hello from V(1).Info()
```
