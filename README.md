## Dagger tutorial

This tutorial showcases two parts:

**1st** Building an application from a remote Git repository and run pipelines through the CLI. 
**2nd** Using the Dagger Go System Development Kit (SDK)

You can find the resources for this tutorial here:
1. The Blog post:
2. The YouTube Video:

## Building an application and running popelines through the CLI

Following the Dagger Documentation: https://docs.dagger.io/cli/389936/run-pipelines-cli
Use the resources in the [following folder.](./one/build.sh)

Using Trivy to scan the project for vulnerabilities with the Trivy-Dagger integration: XXX

## Using the Go SDK

Following the Dagger Documentation: https://docs.dagger.io/sdk/go/
Use the resources in the [following folder.](./one/build.sh)

## Dagger commands

```
dagger --focus=false call build --src ./ export --path ./build-two
```