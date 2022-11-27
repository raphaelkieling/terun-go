<p align="center">
  <img width="200px" src="./logo_primary.png">
</p>
<h1 align="center">Terun</h1>

> It's alpha version using golang. Do not use yet.

Terun is a file transporter, created to help you with big architecture that demands a lot of boilerplate files.

## Install

```
brew install raphaelkieling/terun/terun
```

## Example

Start the `terun.yml`:

```sh
terun init
```

That will create this file:

```yml
commands:
  example:
    args:
      - EntityName
    transports:
      - from: ./controller.template
        to: "./{{.EntityName | underscore | lowercase }}_controller.py"
        name: "Controller file"
        args:
          - CreatedBy
```

Define your template independente of language, in this example i created a file called `controller.template`

> template extension is a recommendation, but feel free to put `anyname.py` or `anyname.js`

```py
## file: controller.template

# author: {{.CreatedBy}}
class {{.EntityName}}Entity{
    constructor(){}
}
```

Run the command `terun make example`, when the global argument request an input you can type `FastPerson` and `My name` for the local argument.

That's the output:

```py
## file: fast_person_controller.py

# author: My name
class FastPersonEntity {
  constructor() {}
}
```

# Advanced

### Recommendations

- Create a folder called `terun` inside your project to allow another devs to reuse the same templates.
- Ever use relative path inside the configuration to avoid problem between environments
- Use a template folder to store your templates. It will keep the `terun` folder tidy.

```diff
my_project/
  main.js
  ...
++terun/
++  terun.yml
++  templates/
++    controller.template
++    service.template
++    repository.template
```

### Template

We are using [golang template engine](https://pkg.go.dev/text/template) that accepts:
- Conditional rendering
- Pipeline functions
- Comments
- Foreach
- ...

### Terun Pipeline Operators

We added some operators inside the buildin golang template engine. That are:

| command      | input     | output    |
| ------------ | --------- | --------- |
| `lowercase`  | `My Name` | `my name` |
| `uppercase`  | `My Name` | `MY NAME` |
| `underscore` | `My Name` | `My_Name` |

Inside the template you can use:

```
{{.Name | lowercase}}
```

Also you are able to use multiples pipelines in the same time.

```
{{.Name | underscore | lowercase}}
```
