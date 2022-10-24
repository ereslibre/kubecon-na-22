package main

import (
	demo "github.com/saschagrunert/demo"
)

func main() {
	// Create a new demo CLI application
	d := demo.New()

	// Register the demo run
	d.Add(example(), "wasmday", "wasmday demo")

	// Run the application, which registers all signal handlers and waits for
	// the app to exit
	d.Run()
}

// example is the single demo run for this application
func example() *demo.Run {
	// A new run contains a title and an optional description
	r := demo.NewRun(
		"Interacting with PHP + mod_wasm",
	)

	r.Step(demo.S(
		"Let's look at a simple PHP script",
	), demo.S(
		"bat ~/php-example/request/index.php",
	))

	r.Step(demo.S(
		"Trigger an HTTP request and inspect the response",
	), demo.S(
		"curl -vvv http://localhost:8000/request/ | jq",
	))

	r.Step(demo.S(
		"Trigger an HTTP request with more parameters and inspect the response",
	), demo.S(
		"curl -vvv",
		"-H \"Other-Header: some-value\"",
		"http://localhost:8000/request/\\?page\\=1\\&offset\\=5 | jq",
	))

	return r
}
