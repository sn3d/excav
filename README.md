<p align="center">
  <h1 align="center">Excav</h1>
  <p align="center">Automatize process of patching repositories in bulk.</p>
  <img align="center" src="https://excav.dev/assets/demo.gif"/>
</p>

---

The goal of excav is not just helping with patching multiple repositories, but 
reducing time spent in general. That means helping with MRs, code review, 
parametrized patching etc.

## How it works

1. create Inventory of repositories you want to patch.

2. define your patch (it could be also reusable parametrized patch)

3. apply patch to selected repositories via `excav apply`

4. push changed to all remote repositories `excav push`

5. check and see merge/pull requests via `excav show`


For more details check the [Quick start](https://excav.dev/quick_start/)

## Get Excav

Please read our [Installation and Configuration](https://excav.dev/installation/)

## Documentation

The full documentation is available [here](https://excav.dev/intro/)

## Bugs & Feature requests

Because it's alpha, you can easily find bugs or you can miss some features.
I will appreciate if you will report bugs and feature requests [here](https://github.com/sn3d/excav/issues)

If you have question of feature request, don't hesitate and visit [Discussion](https://github.com/sn3d/excav/discussions) section.

## Todo

- better installation
- metadata for better exploration of patch parameters
- better code :-)
