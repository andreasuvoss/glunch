# `glunch`

`glunch` is a command line tool for quickly getting access to the lunch menu at work. It was a hassle always having to
find the canteen website, and since most of my work in done in the terminal or an editor nearby my terminal, I thought
it would be convenient to be able to easily list the menu directly from my terminal.

```
$ glunch help
Welcome to glunch - a lunch menu printer written in Go

Available commands:
  glunch            Gets the menu for the current week highlighting today
  glunch help       Shows this help menu
  glunch version    Prints the version of glunch
  glunch w<int>     Gets the menu offset by the number of weeks for example 'glunch w-1'
                    or 'glunch w1' to get the menu of the previous or next week respectively
``
