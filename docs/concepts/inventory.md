# Inventory 
Inventory is simple YAML file called `inventory.yaml` in root of your workspace 
folder. The repositories, you want to patch, are declared here. You can use tags 
for having better control what to patch.

```yaml title="inventory.yaml"
- repository: "/infra-repo/sandbox"
  tags: [ "sandbox", "terraform" ]
- repository: "/infra-repo/production"
  tags: [ "prod", "terraform" ]
```

Not every repository have a default branch `main`. Some legacy repos have `master` as main branch.
This might be determined in inventory via `default_branch`:

```yaml title="inventory.yaml"
- repository: "/infra-repo/sandbox"
  tags: [ "sandbox", "terraform" ]
  default_branch: master
```

### Tags

You can use tags for your repositories. Tags helps you select only subset
of repositories for patching. This is useful when you have GitOps repositories 
for sandbox and production environments and you want to roll-out your patch
on sandboxes first and then on production.

Your repositories will be tagged like:

```yaml title="inventory.yaml"
- repository: "/infra-repo/sandbox-1"
  tags: [ "sandbox" ]
- repository: "/infra-repo/sandbox-2"
  tags: [ "sandbox" ]
- repository: "/infra-repo/production-1"
  tags: [ "prod" ]
```

When use apply with `-tag` flag:

```shell
excav apply -tag sandbox ./hello-patch
```

The patch will be applied only to 2 sandbox repositories from inventory.

### Parameters

Each repository might have own set or parameters they're used in templates.
You can define own parameters via `params` like:

```yaml title="inventory.yaml"
- repository: /org/team2/repo4
  tags: ["prod", "team2"]
  params:
    param1: val1
    param2: val2
```

