# Quick start

First, ensure you have `excav` installed in your system. If not please follow the
instructions in [Installation and Configuration](installation.md) section.

You will need two or more repositores. Create `excav-demorepo-1` and `excav-demorepo-2` repositories in your private or public GitLab or GitHub.

Create inventory file for those repositories in your working directory e.g. `~/excav-demo/inventory.yaml`:

``` yaml title="inventory.yaml"
- repository: "/sn3d/excav-demorepo-1"
- repository: "/sn3d/excav-demorepo-2"
```
    
Then we need create a patch with some simple task. Create a directory 
`demo-patch`.

```sh
mkdocs demo-patch
```

Put a new file `patch.yaml` into this directory with one task:

```yaml title="demo-patch/patch.yaml"
- name: append-hello-world
  append:
    path: README.md
    mode: append-end
    content: "Hello World"
```

This task appends the text `Hello World` at the end of the `README.md` file.

Now we can apply patch to the repositories by running `apply` command.

```sh
excav apply -branch excav/append-text -commit "patching README.md" ./demo-patch
```

The command apply `demo-patch` to all repositories in inventort. Command 
creates commit with name `patching README.md` in own branch `excav/append-text`. 
The changes are local and not pushed yet. You have time to check, if patch is 
applied correctly. For that you can use `diff`.

```sh
excav diff
``` 

If everything looks fine, you can continue with pushing changes.

```sh
excav push
``` 

The command push new branches into remote repositories and create 
Merge Request/Pull Request for a new branch. 

If you want to know if everything is OK, what's the merge/pull request URL etc, 
you can use:

```sh
excav show
```

And if you don't like what patch did, you can discard everything, even MRs via:

```sh
excav discard
```
