# New File
Add a new file into the repository. The task override existing file if it's 
present already. Task creates also parent directories if these are missing.

Example:
```yaml title="patch.yaml"
- name: example-add-readme
  newfile:
    src: readme.txt
    dest: path/to/readme.txt
```

