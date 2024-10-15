# Overview

This repository is a companion to a [blog post](https://www.fermyon.com/blog/api-testing-demo) which walks through how to use Spin's built-in SQLite database, Hurl, and Github Actions to run automated tests against an API built with Spin.

# Usage

## Requirements

- Latest version of [Spin](https://developer.fermyon.com/spin/v2/install)
- Latest version of [TinyGo](https://tinygo.org/getting-started/install/)
- Latest version of [Go](https://go.dev/dl/) supported by TinyGo
  - This can be checked by running `tinygo version`. At the time of this post's publishing, the latest supported version of Go is `1.22`.
- Latest version of [Hurl](https://hurl.dev/docs/installation.html)

## Building, Running, and Testing

To run and test the application, navigate to the root directory of the code. Once there, try the following commands:

### Building the app

```sh
make build
```

### Running the app

```sh
make run
```

### Testing the app

```sh
make test
```

### Stopping the app

```sh
make rm
```
