# Append

This task appends some text or code part into file or files. You can determine 
where to insert the content via `append-begin`, `append-end`.

We can declare content directly in YAML, or for longer snippets, we can
create own template file in patch folder.  

Example how to insert some code fragment from `comment-header.template`
file at the beginning of each `.go` file:

```yaml title="patch.yaml"
- name: example-append-before
  append:
    path: *.go
    mode: append-begin
    template: comment-header.template
```


