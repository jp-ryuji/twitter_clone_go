## Required software

* Go
* PosgreSQL 9.6

## Required libraries

| Type | Stack |
| :----- | :------- |
| Dependency Management | [Go Modules](https://github.com/golang/go/wiki/Modules) |
| ORM | [SQLBoiler](https://github.com/volatiletech/sqlboiler) |

## Development
* Generates a Go ORM from the tables. You should run the commands after modifying the tables.
  ```
  $ bash scripts/sqlboiler.sh
  ```