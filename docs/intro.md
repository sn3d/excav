# Introduction
Excavator (or shortly excav) automatize process of patching repositories in bulk.

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

## How it works

1. create [inventory](concepts/inventory.md) of repositories you want to patch.

2. define your [patch](concepts/patch.md) (it could be also reusable parametrized patch)

3. apply patch to selected repositories via `excav apply`

4. push changed to all remote repositories `excav push`

5. check and see merge/pull requests via `excav show`

Check also our [Quick start](quick_start.md)


