# Remove 
Task removes one or more files. If path points to directory, the whole 
directory will be removed recursively.

```yaml title="patch.yaml"
- name: example-remove-files
  remove:
    files:
      - path/to/file1.txt
      - path/to/dir
```

