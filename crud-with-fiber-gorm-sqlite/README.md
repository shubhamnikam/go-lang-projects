### Crud with go fiber, GORM, Sqlite db

```bash
# to fix - [error] failed to initialize database, got error Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub enable CGO_ENABLE

go env -w CGO_ENABLED=1

install mingw on your system
```

```bash
# sqlite db
    - path is in .env file
```