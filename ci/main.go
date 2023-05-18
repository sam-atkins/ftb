package main

import (
	"context"
	"flag"
	"fmt"
	"os"

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

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	goCache := client.CacheVolume("go")
	source := client.Container().
		From(GoImageVersion).
		WithDirectory("/src", client.Host().
			Directory("."), dagger.ContainerWithDirectoryOpts{
			Exclude: []string{"vendor/"},
		}).
		WithMountedCache("/go", goCache)

	testCmd := []string{"go", "test", "./..."}
	if cover {
		testCmd = append(testCmd, "--cover")
	}
	out, err := source.WithWorkdir("/src").WithExec(testCmd).Stdout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
