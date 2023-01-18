# Replace

Task replace text in file.

Example how to replace text that's matching given reg.exp:


```yaml title="patch.yaml"
- name: example-replace-byregexp
  replace:
    path: example-replace/file1.txt
    regexp: "Hello (.*)"
    replace: "Hello Patch"
```


Example how to replace exact text in file:

```yaml title="patch.yaml"
- name: example-replace-text
  replace:
    path: example-replace/file2.txt
    text: "zdenko-app"
    replace: "patched-app"
```

Example how to replace text in all `*.go` files:


```yaml title="patch.yaml"
- name: example-replace-all-files
  replace:
    path: *.go
    text: "{REPLACE}"
    replace: "Hello Excav"
```

Example how to replace some piece of text by content from template file:

```yaml title="patch.yaml"
- name: replace-template
  replace:
    path: example-replace/file3.txt
	 text: "TODO: text here"
	 template: template.txt
```

