## jumpo (ju)

Jump list for your shell. It is a small binary that can be used to quickly move
between directories for very fast workflow 🚀. It works on a very simple concept
to use simple persistent key-value pair (JSON file on disk 😛) and prints the
valid `cd` command on the `stdout` which will be later consumed by the bash
function to `cd` into that directory.

## Demo

<img src="./assets/output.gif">

## Usage

Firstly put this function in your shell config (`.bashrc` or `.zshrc`)
```bash
ju() {
    `jumpo-mac $*`
}
```

Then after sourcing the shell config again or restarting the terminal. You will
be able to see the function.

```bash
functions | grep 'ju ()'
```

```bash
# Last argument is the prefix string, which will be used for jumping
# Will use the directoy of execution as the location of jump.
ju -add no

# To move that directory
ju no
```

Available options are as follows:

```bash
# To add current directory to jump list with prefix
# Or use the short-hand command -a
ju -add r1

# List the jumplist
# Or use the short-hand command -l
ju -list

# Remove prefix from jumplist
# Or use the short-hand command -r
ju -remove r1

# Move to the jump list
ju r1
```

## Building

There is a simple Makefile present in the project which will build binaries for
different platforms.

```bash
make
```
