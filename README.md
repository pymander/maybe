# maybe
A simple command line utility that succeeds sometimes.

## Installation

With Go installed:

```bash
go install github.com/pymander/maybe@latest
```

Or build from source:

```bash
git clone https://github.com/pymander/maybe
cd maybe
go build
```

Binary will be installed to `$GOPATH/bin` when using `go install`, or in the current directory when using `go build`.

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
