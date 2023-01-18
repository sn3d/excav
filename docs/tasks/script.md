# Script

This task executes any custom script over the patching repository. 
This task is for you, if you need some advanced patching. You can use it 
for python patching script. Excav pass the repository path into the 
script via `{{.RepositoryDir}}` placeholder.

Example executes `script.py` python script and pass the absolute path to 
repository as first argument.

```yaml title="patch.yaml"
- name: example-call-script
  script:
    - python script.py {{.RepositoryDir}}
```
