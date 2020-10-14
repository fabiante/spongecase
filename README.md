# sPoNgeCaSE

> Inspired by: https://knowyourmeme.com/memes/mocking-spongebob

This is both a Go module and a command line tool converts strings into "Sponge Case".

I actually don't know if it is called like that but suit yourself 😉

This tool will convert strings like this ...

```text
This will definitely build my reputation as a software developer
```

into this ...

```text
THIs Will Definitely bUIld mY RepUTaTiOn AS a soFtWaRe deVelOPER
```

## Installation

You can either download prebuild executables from the [Releases tab](https://github.com/FabianTe/spongecase/releases) or from the the [AUR](https://aur.archlinux.org/packages/spongecase-git/)

## Usage

#### CLI - Piping string into STDIN

```shell script
echo "This will definitely build my reputation as a software developer" | spongecase
```

#### CLI - Passing string as argument

```shell script
spongecase "This will definitely build my reputation as a software developer"
```

#### Importing as Go module

```shell script
go get github.com/FabianTe/spongecase
```

```go
package main

import (
	"fmt"
	"github.com/FabianTe/spongecase"
)

func main() {
	fmt.Println(spongecase.ApplyStr("This will definitely build my reputation as a software developer"))
}
```