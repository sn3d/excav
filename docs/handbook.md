# Handbook

In this handbook you will find detailed informations how to define inventory of
repositories,how to define patches etc.

## Quick start

Before you start, create `excav-demorepo-1` and `excav-demorepo-2` repositories in your private or 
public GitLab or GitHub. Also be sure your `excav` is installed and configured 
(see [Installation and Configuration](../README.md#installation-and-configuration)). 

Create inventory file in some workspace directory e.g. `~/excav-demo/inventory.yaml`:

```
- repository: "/sn3d/excav-demorepo-1"
  tags: [ "demorepo" ]
- repository: "/sn3d/excav-demorepo-2"
  tags: [ "demorepo" ]
```
    
Then we need create some patch with some simple task e.g. 
`~/excav-demo/append-text-to-tile/patch.yaml`:

```
- name: append-hello-world
  append:
    path: README.md
    mode: append-end
    content: "Hello World"
```

This task appends the text `Hello World` at the end of the `README.md` file.

Now we can apply patch to the repositories by running `apply` command.

```
excav apply -tag demorepo -branch excav/append-text -commit "append text" ./append-text-to-file
```

The command apply `append-text-to-file` to all repositories with `demorepo` tag. 
Command creates commit `apply patch` in own branch `excav/append-text`. The 
changes are not pushed yet. You have time to check how patch is applied. For 
that you can use `diff`.

```
excav diff
``` 

If everything looks fine, you can continue with pushing changes.

```
excav push
``` 

The command push new branches into remote repositories and create 
Merge Request/Pull Request for a new branch. 

If you want to know if everything is OK, what's the MR URL etc, you can use
```
excav show
```

And if you don't like what patch did, you can discard everything, even MRs via:

```
excav discard
```

## Inventory 
Inventory is simple YAML file called `inventory.yaml` in root of your workspace 
folder. The repositories, you want to patch, are declared here. You can use tags 
for having better control what to patch.

```
- repository: "/infra-repo/sandbox"
  tags: [ "sandbox", "terraform" ]
- repository: "/infra-repo/production"
  tags: [ "prod", "terraform" ]
```

Not every repository have a default branch `main`. Some legacy repos have `master` as main branch.
This might be determined in inventory via `default_branch`:

```
- repository: "/infra-repo/sandbox"
  tags: [ "sandbox", "terraform" ]
  default_branch: master
```

### Parameters

Each repository might have own set or parameters they're used in templates.
You can define own parameters via `params` like:

```
- repository: /org/team2/repo4
  tags: ["prod", "team2"]
  params:
    param1: val1
    param2: val2
```


## Patch

Excav organize patches in own folders same as Ansible organize Roles. It's because 
some patches might have assets like templates or source code files. Every patch 
folder has `patch.yaml` on root, where we declare [tasks](#tasks).

Example of `my-patch.yaml` patch with single task

```
- name: append-hello-world
  append:
    path: README.md
    mode: append-end
    content: "Hello World"
```

## Tasks

Task is simple single step what to do. Excav executes the task for each 
repository. Excav supports following types of tasks you can use in your patch.

### Append
 
Append appends some text or code part into file or files. You can determine 
where to insert the content via `append-begin`, `append-end`.

We can declare content directly in YAML, or for longer snippets, we can
create own template file in patch folder.  

Example how to insert some code fragment from `comment-header.template`
file at the beginning of each `.go` file:

```
- name: example-append-before
  append:
    path: *.go
    mode: append-begin
    template: comment-header.template
```

### Put
Put puts some text at the specific part of the file. This task looks for
anchor and put the content immediately after this anchor.

Let's imagine you have some code with anchor `+excav:put:here`:

```
func main() {
   ...
   // +excav:puthere
   ...
}
```

And you want to put some fragment from `code.template` after this anchor:

```
- name: put-snippet
  put:
    path: "main.go"
    anchor: "\\/\\/ \\+excav:puthere"
    content: "x := x + 1"
```

The result will be:

```
func main() {
   ...
   // +excav:puthere
   x := x + 1
   ...
}
```

Notice the code is auto-aligned with anchor by intends.

You can also extract some parameter from anchor and use it in your 
content/template. Let's imagine the code `main.go` with anchor:

```
func main() {
   ...
   // +excav:puthere:1
   ...
   // +excav:puthere:2
}
```

And we have put task where we can extract the value from anchor

```
- name: put-snippet
  put:
    path: "main.go"
    anchor: "\\/\\/ \\+excav:puthere:(.*)"
    content: "// Hello {{index .groups 0}}"
```

The output will be:

```
func main() {
   ...
   // +excav:puthere:1
   // Hello 1
   ...
   // +excav:puthere:2
   // Hello 2
}

```

### New File
Add a new file into the repository. The task override existing file if it's 
present already. Task creates also parent directories if these are missing.

Example:
```
- name: example-add-readme
  newfile:
    src: readme.txt
    dest: path/to/readme.txt
```

### Remove 
Task removes one or more files. If path points to directory, the whole 
directory will be removed recursively.

```
- name: example-remove-files
  remove:
    files:
      - path/to/file1.txt
      - path/to/dir
```

### Replace

Task replace text in file.

Example how to replace text that's matching given reg.exp:
```
- name: example-replace-byregexp
  replace:
    path: example-replace/file1.txt
    regexp: "Hello (.*)"
    replace: "Hello Patch"
```


Example how to replace exact text in file:
```
- name: example-replace-text
  replace:
    path: example-replace/file2.txt
    text: "zdenko-app"
    replace: "patched-app"
```

Example how to replace text in all '.go' files:
```
- name: example-replace-all-files
  replace:
    path: *.go
    text: "{REPLACE}"
    replace: "Hello Excav"
```

Example how to replace some piece of text by content from template file:
```
- name: replace-template
  replace:
    path: example-replace/file3.txt
	 text: "TODO: text here"
	 template: template.txt
```

### Script
Task executes any script over the patching repository. This task is for you, 
if you need some advanced patching. You can use it for python patching 
script. Excav pass the repository path into the script via {{.RepositoryDir}} 
placeholder.

Example executes `script.py` python script and pass the absolute path to 
repository as first argument.
```
- name: example-call-script
  script:
    - python script.py {{.RepositoryDir}}
```
