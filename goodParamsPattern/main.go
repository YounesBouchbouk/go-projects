package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxCount int
	id       string
	tls      bool
}

func defaultOpts() Opts {
	return Opts{
		maxCount: 10,
		id:       "default",
		tls:      false,
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withCostumTLS(tls bool) OptFunc {
	return func(ops *Opts) {
		ops.tls = tls
	}
}

func withMaxCount(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxCount = n
	}
}

func withId(i string) OptFunc {
	return func(opts *Opts) {
		opts.id = i
	}
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {

	o := defaultOpts()

	for _, fn := range opts {

		fn(&o)

	}
	return &Server{
		Opts: o,
	}
}

// func testfunc(params ...string) {
// 	fmt.Printf("%+v\n", params[0])

// }

func main() {
	// testfunc("hello", "bonjour", "salamO3alikom")

	//new we can user
	s := newServer(withTLS, withMaxCount(44), withId("hellp"))
	// or
	// s := newServer(withId("hellp")) or s := newServer(withTLS)

	fmt.Printf("%+v\n", s)
}
