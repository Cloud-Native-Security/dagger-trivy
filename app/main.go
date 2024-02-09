package main

import (
	"context"
	"fmt"
)

type Example struct{}

const (
	trivyImageTag = "0.49.1" // semver tag or "latest"
)

// example usage: "dagger call container-echo --string-arg yo"
func (m *Example) ContainerEcho(stringArg string) *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// example usage: "dagger call grep-dir --directory-arg . --pattern GrepDir"
func (m *Example) GrepDir(ctx context.Context, directoryArg *Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

func (m *Example) Build(ctx context.Context, src *Directory) *Directory {

	// get `golang` image
	golang := dag.Container().From("golang:latest")

	// mount cloned repository into `golang` image
	golang = golang.WithDirectory("/src", src).WithWorkdir("/src")

	// define the application build command
	path := "build/"
	golang = golang.WithExec([]string{"env", "GOOS=darwin", "GOARCH=arm64", "go", "build", "-o", path, "./src/hello.go"})

	scanResult, err := dag.Trivy().ScanContainer(ctx, golang, TrivyScanContainerOpts{
		TrivyImageTag: trivyImageTag,
		Severity:      "HIGH,CRITICAL",
		Format:        "table",
		ExitCode:      0,
	})

	fmt.Println(scanResult)

	if err != nil {
		return nil
	}

	// get reference to build output directory in container
	outputDir := golang.Directory(path)

	return outputDir
}
