# Patch

Excav organize patches in own folders same as Ansible organize Roles. It's because 
some patches might have assets like templates or source code files. Every patch 
folder has `patch.yaml` on root, where we declare [tasks](../tasks/index.md).

Example of patch with single [append task](../tasks/append.md):

```yaml title="demo-patch/patch.yaml"
- name: append-hello-world
  append:
    path: README.md
    mode: append-end
    content: "Hello World"
```

