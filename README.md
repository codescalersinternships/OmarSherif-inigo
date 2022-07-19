# OmarSherif-inigo

## if you want to test
```
go test
```

## if you want to run

```
go run !(*_test).go
```


## All Functions

* `LoadFromString() ` : Loads a string into a `Parser` object to generate a map of sections .
* `LoadFromFile()` : Loads a file into a `Parser` object to generate a map of sections .
* `GetSectionNames()` list of all section names
* `GetSections()` serialize convert into a dictionary/map { section_name {key1: val1, key2, val2} ...}
* `Get(section_name, key)` gets the value of key key in section section_name
* `Set(section_name, key, value)` sets a key in section section_name to value value
* `String()`: serialize convert into a string
* `SaveToFile()`: saves the current state of the `Parser` object by applying on it String() to a file
