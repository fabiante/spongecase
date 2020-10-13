# Sponge Case

> Inspired by: https://knowyourmeme.com/memes/mocking-spongebob

This is both a go module and a command line tool converts strings into "Sponge Case".

I actually don't know if it is called like that but suit yourself 😉

This tool will convert strings like this ...

```text
This will definately build my reputation as a software developer
```

into this ...

```text
THIs Will DEfinateLy bUIld mY RepUTaTiOn AS a soFtWaRe deVelOPER
```

## Usage

#### CLI - Piping string into STDIN

```shell script
echo "This will definately build my reputation as a software developer" | spongecase
```

#### CLI - Passing string as argument

```shell script
spongecase "This will definately build my reputation as a software developer"
```