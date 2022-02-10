package main

import (
	"flag"

	"github.com/golang/glog"
)

type Level int32

// Ensure that Level implements the interface flag.Value
var _ flag.Value = (*Level)(nil)

var verbosity Level

// Dummy implementations
func (l *Level) Get() interface{}     { return nil }
func (l *Level) Set(val string) error { return nil }
func (l *Level) String() string       { return "" }

func init() {
	/*
	 * Trying to declare an option 'v' causes a Panic at runtime, because the
	 * glog package has already registered that, and a few other options, in its
	 * init() method.
	 */
	flag.Var(&verbosity, "v", "Emit verbose execution progress")
}

func main() {
	// Parse the command-line options
	flag.Parse()

	glog.V(1).Info("Hello from V(1).Info()")
}
