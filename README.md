# maybe
A simple command line utility that succeeds sometimes.

## Installation

Since there are no binaries being distributed at this time, you will need to [install Go](https://go.dev/dl/), at least version 1.20. You might as well just get the newest one, you won't regret it.

With Go installed, you can do the remote install thing with just one step:

```bash
go install github.com/pymander/maybe@latest
```

Or go slow and steady and build from source:

```bash
git clone https://github.com/pymander/maybe
cd maybe
go build
```

The `maybe` binary will be installed to `$GOPATH/bin` when using `go install`, or in the current directory when using `go build`.

## Usage

```
  -chance int
        Percent chance from 1 to 99 (default 50)
  -verbose
        Be noisy
```

The percent chance is used to determine whether maybe exits with a 0 or a 1. This lets you create, for example, cron jobs that have a chance of running, but won't run all the time. You can do something like this:

```bash
maybe -chance 75 && echo Hello world
```

Who knows what crazy things you only want to happen part of the time?
