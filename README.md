A simple tool to conceal your confidential string in a public environment then reveal it in program.

### Usage

1. In your golang source fie:

```
$ go get -u github.com/gogoods/faker
```

```
package main

import "github.com/gogoods/faker"

func main(){

    fakepass := "xxxxxx"
    realpass := faker.Reveal(fakepass)
}
```

2. Command line:

```
$ go install -u github.com/gogoods/faker
```

```
$ faker -h
$ faker conceal xxxxx
$ faker reveal xxxxx
```