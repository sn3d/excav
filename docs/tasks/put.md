# Put
Put puts some text at the specific part of the file. This task looks for
anchor and put the content immediately after this anchor.

Let's imagine you have some code with anchor `+excav:put:here`:

```go title="main.go"
func main() {
   ...
   // +excav:puthere
   ...
}
```

And you want to put some fragment from `code.template` after this anchor:

```yaml title="patch.go"
- name: put-snippet
  put:
    path: "main.go"
    anchor: "\\/\\/ \\+excav:puthere"
    content: "x := x + 1"
```

The result will be:

```go title="main.go"
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

```go title="main.go"
func main() {
   ...
   // +excav:puthere:1
   ...
   // +excav:puthere:2
}
```

And we have put task where we can extract the value from anchor

```yaml title="patch.yaml"
- name: put-snippet
  put:
    path: "main.go"
    anchor: "\\/\\/ \\+excav:puthere:(.*)"
    content: "// Hello {{index .groups 0}}"
```

The output will be:

```go title="main.go"
func main() {
   ...
   // +excav:puthere:1
   // Hello 1
   ...
   // +excav:puthere:2
   // Hello 2
}

```


