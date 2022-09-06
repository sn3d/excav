# Excav

Excavator (or shortly excav) automatize process of patching repositories in bulk.

![](./docs/demo.gif)

## How it works

1. create [inventory](docs/handbook.md#inventory) of repositories you want to patch.

2. define your [patch](docs/handbook.md#patch) (it could be also reusable parametrized patch)

3. apply patch to selected repositories via `excav apply`

4. push changed to all remote repositories `excav push`

5. check and see merge/pull requests via `excav show`

For more details check the [Quick start](docs/handbook.md#quick-start)

## Motivation

With raise of GitOps and microservices, we're facing to many repositories. Many of 
them sharing common concepts, and it's easy to patch them in bulk. The problem is, 
we need to write own scripts they're going through those repositories and apply some 
simple operations. 

The goal of excav is not just helping with patching itself but reducing time spent in 
general. That means helping with MRs, code review etc.

Of course, not every patch can be applied in bulk. The goal of this tool is to
help with those, they're easily reproducible. You need to consider if you're 
able to patch them easily, or you need some very specific patching.

## Installation and configuration

Currently, `excav` is not available in any package system like Homebrew etc. But 
it's simple binary file which can be easily installed with those two methods:

### curl/wget install (macOS & Linux)

You can install excav with curl or get with those one-liners:

```
curl -s https://installme.sh/sn3d/excav | sh
```

```
wget -q -O - https://installme.sh/sn3d/excav | sh
```

### Manual install (mainly for Windows)

Download the correct binary for your platform from [project's GitHub](https://github.com/sn3d/excav/releases/latest).
Uncompress the binary to you `PATH`.


### Configuration 

After installation is done, you can configure it by running the `init`. You will 
be prompted few questions about GitLab etc.

```
excav init 
``` 

Initialization process creates `~/.config/excav/config.yaml` for you. You can
do any further configuration changes in this file.

## Documentation

The detailed documentation is available [here](docs/handbook.md)

## Bugs & Feature requests

Because it's alpha, you can easily find bugs or you can miss some features.
I will appreciate if you will report bugs and feature requests [here](https://github.com/sn3d/excav/issues)

## Todo

- better installation
- support for GitHub (It's already implemented but not tested well)
- metadata for better exploration of patch parameters

