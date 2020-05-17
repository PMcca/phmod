# phmod

phmod is a simple translator for `chmod` written in Go (1.14). Given an octal (771) or symbolic (rwxr-x--x) representation of a chmod input, phmod will produce its opposite with the option of some explanation using the `-l` flag.

## Install
Run `GO111MODULE=on go get github.com/PMcca/phmod` to install the package to your $GOPATH/bin directory. Next run `ln -s $GOPATH/bin/phmod /usr/local/bin/phmod` to run it via your $PATH.

## Usage
```
phmod [OPTION] <args>

Options:
--long, -l     Prints the result in verbose, human-readable format.
--help, -h     show help
--version, -v  print the version
```
**NOTE**
To give an input such as `--xrwxrwx` (that is, the argument starting with `-`), precede the argument with `--`. e.g.

`phmod -- --xrwxrwx`

Output: `177`

### Example 
`phmod 571`
Output: `r-xrwx--x`

`phmod -l r-xrwx--x`
Output: 
```
User Modes:
RWX
101 = 5

Group Modes:
RWX
111 = 7

Other Modes:
RWX
001 = 1
```
