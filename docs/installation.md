# Installation and configuration

## Installation

Currently, `excav` is not available in any package system like Homebrew etc. But 
it's simple binary file which can be easily installed with those two methods:

### MacOS & Linux

You can install excav with curl or get with those one-liners:

```sh
curl -s https://installme.sh/sn3d/excav | sh
```

```sh
wget -q -O - https://installme.sh/sn3d/excav | sh
```

### Windows

Download the correct binary for your platform from [project's GitHub](https://github.com/sn3d/excav/releases/latest).
Uncompress the binary to you `PATH`.


## Configuration 

After installation is done, you can configure it by running the `init`. You will 
be prompted few questions about your GitHub/GitLab etc.

```sh
excav init 
``` 

Initialization process creates `~/.config/excav/config.yaml` for you. You can
do any further configuration changes in this file.

