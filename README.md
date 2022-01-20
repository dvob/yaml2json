# yaml2json
Convert YAML files to JSON

Usage:
```
yaml2json FILE [FILE1...]
```

## Example
* `foo.yaml`
```
foo:
  bar: bla
  baz: 1
```
```
$ yaml2json foo.yaml
{
  "foo": {
    "bar": "bla",
    "baz": 1
  }
}
```

* From standard input
```
cat bla.yaml | yaml2json -
```
