## Dagger tutorial

This tutorial showcases two parts:

**1st** Building an application from a remote Git repository and running pipelines through the CLI. 
**2nd** Using the Dagger Go System Development Kit (SDK)

Related YouTube Videos:
1. Dagger Live Stream: https://www.youtube.com/live/Rc0NCpc2YSc?si=L7IYIMpXubdEbsZL
2. Dagger Community Call: https://youtu.be/s9BBgp_tOf0?si=GyJ05kL98poc2xHz

## Building an application and running pipelines through the CLI

Following the Dagger Documentation: https://docs.dagger.io/cli/389936/run-pipelines-cli
Use the resources in the [following folder.](./one/build.sh)
## Using the Go SDK

Following the Dagger Documentation: https://docs.dagger.io/sdk/go/
Use the resources in the [following folder.](./one/build.sh)

## Dagger commands

```
dagger --focus=false call build --src ./ export --path ./build-two
```
