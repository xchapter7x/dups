# dups [![wercker status](https://app.wercker.com/status/abc70f396f16efc5f27747175a33f956/s/master "wercker status")](https://app.wercker.com/project/byKey/abc70f396f16efc5f27747175a33f956)
find duplicates in a list of files

Install:

```
go get github.com/xchapter7x/dups
```

Sample Usage:

```
#search for files and have grup show dups
$ find /repo/path -name "*_test.go" | dups
```


Summary:

defaults to finding duplicates in the set of files given, using a 10 line block comparision.
