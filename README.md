## Required software

* Go
* PosgreSQL 9.6

## Required libraries

| Type | Stack |
| :----- | :------- |
| Dependency Management | [Go Modules](https://github.com/golang/go/wiki/Modules) |
| ORM | [SQLBoiler](https://github.com/volatiletech/sqlboiler) |

## Development
* Set the following environment variables in `.envrc` or somewhere to use `Go Modules`.
  ```
  export GO111MODULE=on
  ```

* Generates a Go ORM from the tables. You should run the commands after modifying the tables.
  ```
  $ bash scripts/sqlboiler.sh
  ```