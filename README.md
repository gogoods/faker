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

        conceal <content> [prefix]     Encode a string and add a faker prefix.
        reveal  <content> [prefix]     Decode a string if it contains a faker prefix.        
```


```
$ faker conceal xxxxx
$ faker reveal xxxxx

$ faker conceal xxxxx  @
$ faker reveal xxxxx  @
```