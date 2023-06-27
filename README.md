# gbms

General background management system

## Reqired

### Ubuntu

```bash
apt-get install postgresql-client postgresql-client-common
```

## Test

> #### Notice
>> Please check pg username and password in Makefile are correct

```bash
make test
```


## 
```bash
Go mod tidy
go build
./gbms > /dev/null 2>&1 &
```
