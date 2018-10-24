
# CourseHub
> This is material used to explore testing in Go language.

Course service like coursera, udemy, udacity, etc.



## Development - How to

- Put the source code under `$GOPATH/src/github.com/uudashr/coursehub`
- Install all required tools
  ```
  make prepare-all
  ```
- Download vendor dependencies
  ```
  make vendor
  ```


## Test

### Unit Test

```
make test
```

### Repository Test

Run `mysql` test. Test for `mysql` package consider as integration test.

```
# start MySQL on docker (ctrl+c if the log no longer rolls)
make docker-mysql-up

# run the test
make test-mysql

# stop the MySQL
make docker-mysql-down
```

### Test all
Run all test (include the integration test)

```
# start MySQL on docker (ctrl+c if the log no longer rolls)
make docker-mysql-up

# run the test
make test-all

# stop the MySQL
make docker-mysql-down
```


## Reference

- https://github.com/stretchr/testify#mock-package
- https://github.com/vektra/mockery
- https://github.com/golang-migrate/migrate

