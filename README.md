<p align="center">
  <img width="100px" src="./logo.png">
</p>
<h1 align="center">Terun</h1>

> It's alpha version using golang. Do not use yet.

Terun is a file transporter, created to help you with big architecture that demands a lot of boilerplate files.

## Install

```
brew install terun
```

## Example

Start the `terun.yml`:

```sh
terun init
```

Define your template independente of language:

```javascript
// file: from.template
class {{.EntityName | capitalize}}Entity{
    constructor(){}
}
```

Run on terminal `terun make example`:

```javascript
// file: person.py
class PersonEntity {
  constructor() {}
}
```
