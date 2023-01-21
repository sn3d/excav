# File
Add a new file into the repository or override existing one. Task creates also 
parent directories if these are missing.

Example:
```yaml title="patch.yaml"
- name: readme
  newfile:
    src: readme.txt
    dest: path/to/readme.txt
```

