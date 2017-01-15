# gotpl - CLI tool for Golang templates

Command line tool that compiles Golang
[templates](http://golang.org/pkg/text/template/) with values from YAML files.

## Install

    go get github.com/jgensler8/gotpl

## Usage

Contents of `template`:

```
Hello, {{.name}}. A friendly reminder that {{.note.json}}
```

Contents of `input.yaml`

```
---
name: <yourname>
note: {
  "json": "json is valid yaml! :)"
}
```

```
$ ./gotpl -data examples/input/input.yaml -template examples/input/template.tmpl
Hello, <yourname>. A friendly reminder that json is valid yaml! :)
```


## Dockerfile

```
GOOS=linux go build
docker build -t gotpl .
docker run -v "$PWD/examples/input/input.yaml:/input/data" -v "$PWD/examples/input/template.tmpl:/input/template"  gotpl
```
