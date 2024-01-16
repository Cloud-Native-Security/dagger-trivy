package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory("./src")
	mod := client.Host().File("./go.mod")
	sum := client.Host().File("./go.sum")

	// get `golang` image
	golang := client.Container().From("golang:latest")

	// mount cloned repository into `golang` image
	golang = golang.WithDirectory("/src", src).WithWorkdir("/src")
	golang = golang.WithFile("go.mod", mod)
	golang = golang.WithFile("go.sum", sum)

	// define the application build command
	path := "build/"
	golang = golang.WithExec([]string{"env", "GOOS=darwin", "GOARCH=arm64", "go", "build", "-o", path})

	// get reference to build output directory in container
	output := golang.Directory(path)

	// write contents of container build/ directory to the host
	_, err = output.Export(ctx, path)
	if err != nil {
		return err
	}

	return nil
}
