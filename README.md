Steps to Reproduce
==================

```
$ git clone https://github.com/gurjeet/glog_fix_test

$ cd glog_fix_test

$ go build -o main && ./main
./main flag redefined: v
panic: ./main flag redefined: v
... stack trace follows
```

How to Test the Fix
===================

```
# Inside the same directory, clone the glog repo
$ git clone https://github.com/golang/glog.git
$ pushd glog

# Switch to the repo and branch that has the fix
$ git remote add gurjeet https://github.com/gurjeet/glog
$ git checkout gurjeet/option_to_not_pollute_flags
# Alternatively, apply the patch included in this repository
#$ patch -p1 < ../bugfix.patch

$ popd

# Tell the compiler to use our local version of glog
$ go mod edit -replace=github.com/golang/glog=./glog

# Ask the linker to assign a value to the global variable FlagPrefix
$ go build -o main -ldflags "-X github.com/golang/glog.FlagPrefix=glog:"
$ ./main -glog:logtostderr -glog:v 1 -v somevalue

I0210 11:17:28.096799   72018 main.go:34] Hello from V(1).Info()

# Ask the linker to assign a different value to the global variable FlagPrefix
$ go build -o main -ldflags "-X github.com/golang/glog.FlagPrefix=abc"
$ ./main -abclogtostderr -abcv 1 -v somevalue
```
