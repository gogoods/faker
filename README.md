A simple tool to conceal your confidential string in a public environment then reveal it in program.

### Usage

```
$ go get -u github.com/gogoods/faker
```

1. Command line:

```
$ go install github.com/gogoods/faker
```

```
$ faker -h


        ▄████  ██   █  █▀ ▄███▄   █▄▄▄▄
        █▀   ▀ █ █  █▄█   █▀   ▀  █  ▄▀
        █▀▀    █▄▄█ █▀▄   ██▄▄    █▀▀▌
        █      █  █ █  █  █▄   ▄▀ █  █
         █        █   █   ▀███▀     █
          ▀      █   ▀             ▀
                ▀

Faker commands:

        enc <content>      Encode a string.
        dec <content>      Decode a string.

        conceal <content> [prefix]     Encode a string and add a faker prefix. Default prefix is '~'.
        reveal  <content> [prefix]     Decode a string if it contains a faker prefix. Default prefix is '~'.    
```


```
$ faker conceal xxxxx
~1pFwwweIL5EXa3GiWvsyZg==
$ faker reveal ~1pFwwweIL5EXa3GiWvsyZg==
xxxxx

$ faker conceal xxxxx  ^
^1pFwwweIL5EXa3GiWvsyZg==
$ faker reveal xxxxx  @
xxxxx
```

2. In your golang source fie:

```
package main

import "github.com/gogoods/faker"

func main(){

    fakepass := "~1pFwwweIL5EXa3GiWvsyZg=="
    realpass := faker.Reveal(fakepass)   //xxxxx
}
```