# Scribe: A tool that helps you write better git commit messages

You may be one of those people who write commit messages like *"fix"*, *"fix2"*, or *"zzzzz"*. If so, this tool helps you write them better by forcing a more structured message.

This project is written in Go and uses [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the terminal UI.

## So how does it work ?

When you run `git commit`, Git calls this program through `core.editor` to handle the commit message. It forces you to specify:

- Commit Type
- Commit Title
- Commit Description (optional)

---
    Note that the commit types are enumerated.
---

## How to install it

### On Linux

Clone this repository by running the following command:

    git clone https://github.com/LyEXS/Scribe

Then build the source code by running:

    go build -o scribe

Move the binary to a folder in your `PATH`:

    sudo mv scribe /usr/local/bin/

Finally, set Git to call your program:

    git config --global core.editor "scribe"

### On Windows

Clone this repository by running the following command:

    git clone https://github.com/LyEXS/Scribe

Then build the source code by running:

    go build -o scribe.exe

Move the executable to a folder that is in your `PATH`:

    move scribe.exe C:\Program Files\scribe\

Finally, set Git to call your program:

    git config --global core.editor "scribe"

### On Mac

Clone this repository by running the following command:

    git clone https://github.com/LyEXS/Scribe

Then build the source code by running:

    go build -o scribe

Move the binary to a folder in your `PATH`:

    sudo mv scribe /usr/local/bin/

Finally, set Git to call your program:

    git config --global core.editor "scribe"
