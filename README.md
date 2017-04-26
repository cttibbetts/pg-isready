# pg-isready

Simple app that waits for a postgres connection before exiting.

### Environment Variables
```
POSTGRES_USER - Username
POSTGRES_PASS - Password
POSTGRES_HOST - Hostname or IP
POSTGRES_PORT - Port Number
POSTGRES_DB   - Database name
```

# Vendoring
`pg-go` uses [govendor](https://github.com/kardianos/govendor) for vendoring.
To pull the vendor files, run:
```
go get -u github.com/kardianos/govendor
govendor sync
```

