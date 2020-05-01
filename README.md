# phmod

phmod is a simple translator for `chmod` written in Go (1.14). Given an octal (771) or symbolic (rwxr-x--x) representation of a chmod input, phmod will produce its opposite with the option of some explanation using the `-l` flag.

## Install
Run (some command) to install the package to your $GOPATH/bin directory. Assuming your $PATH is set up to include your $GOPATH / $GOROOT bin directory, the program should be ready to run upon restarting your terminal (`source ~/.bashrc` usually)

## Usage
```
phmod [OPTION] <args>

Options:
--long, -l     Prints the result in verbose, human-readable format.
--help, -h     show help
--version, -v  print the version
```
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