glog Library Pollutes Global Flags List
=======================================

This problem has been reported quite a few times, over the years; see it
reported at [golang-nuts][] mailing list, and [cue lang issue][].

[golang-nuts]: https://groups.google.com/g/golang-nuts/c/vj8ozVqemnQ
[cue lang issue]: https://github.com/cue-lang/cue/issues/1199

The problem is that, that glog package registers some flags in its `init()`
function. The list of registered flags also includes the `-v`  flag, which is
usually used by developers to either control verbosity of their code-execution,
or to show the software version. It's notable that all the compaints are
regarding the `-v` flag, and none of the other flags, since those other ones are
unlikely be used by any other developer.

The proposed fix allows the user of the glog library to change/prefix glog's
flags' names, so that they will not conflict with any flags that they wish to
use.

This approach to the problem has a few advantages, compared to other approaches
like, disabling all the flags in glog.

.) The default behaviour of the glog library is unchanged. So the current users of the library will not be affected.

.) Any new users who wish to use the `-v`, or other glog-occupied flag, can do so at build time.

.) The new users can still use the glog features/flags, albeit with a prefix.

.) We are not enforcing some specific prefix, which may also conflict.

.) The `--help` flag, correctly reports the changed/prefixd flag names.

```
$ ./main --help
Usage of ./main:

  ... other glog: prefixed flags ...

  -glog:v value
        log level for V logs
  -glog:vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
  -v value
        Emit verbose execution progress

```

Steps to Reproduce
==================

```
$ git clone https://github.com/gurjeet/glog_fix_test

$ cd glog_fix_test

$ go build -o main && ./main
./main flag redefined: v
panic: ./main flag redefined: v
... stack trace follows ...
```

Steps to Test the Fix
=====================

```
# Inside the same directory, clone the glog repo
$ git clone https://github.com/golang/glog.git
$ pushd glog

# Switch to the repo and branch that has the fix
$ git remote add gurjeet https://github.com/gurjeet/glog
$ git fetch --all
$ git checkout option_to_not_pollute_flags
# Alternatively, apply the patch included in this repository
#$ patch -p1 < ../bugfix.patch

$ popd

# Tell the compiler to use our local/pathced version of glog
$ go mod edit -replace=github.com/golang/glog=./glog

# Ask the linker to assign a value to the global variable FlagPrefix
$ go build -o main -ldflags "-X github.com/golang/glog.FlagPrefix=glog:"
$ ./main -glog:logtostderr -glog:v 1 -v somevalue

I0210 11:17:28.096799   72018 main.go:34] Hello from V(1).Info()

# Ask the linker to assign a different value to the global variable FlagPrefix
$ go build -o main -ldflags "-X github.com/golang/glog.FlagPrefix=abc"
$ ./main -abclogtostderr -abcv 1 -v somevalue
```
