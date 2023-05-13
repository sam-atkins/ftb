package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
)

const (
	GoImageVersion = "golang:1.20.4-bullseye"
)

var (
	cover bool
)

func main() {
	flag.BoolVar(&cover, "cover", false, "Run tests with coverage")
	flag.Parse()
	RunTests()
}

func RunTests() {
	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// create a cache volume
	goCache := client.CacheVolume("go")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	cacheDir := filepath.Join(homeDir, "go/pkg/mod")

	// use a golang container
	// mount the source code directory on the host
	// at /src in the container
	source := client.Container().
		From(GoImageVersion).
		WithDirectory("/src", client.Host().
			Directory("."), dagger.ContainerWithDirectoryOpts{
			Exclude: []string{"vendor/"},
		}).
		WithMountedCache(cacheDir, goCache)

	// set the working directory in the container
	runner := source.WithWorkdir("/src")

	// run application tests
	var testCmd []string
	if cover {
		testCmd = []string{"go", "test", "./...", "--cover"}
	} else {
		testCmd = []string{"go", "test", "./..."}
	}
	out, err := runner.WithExec(testCmd).Stdout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
