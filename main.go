package main

import "glunch/cmd"
import "os"
import "fmt"
import "text/tabwriter"
import "strconv"

func main() {
	args := os.Args

	if len(args) == 1 {
		cmd.GetMenu(0)
	} else if len(args) == 2 {

		arg := args[1]

		if arg[0:1] == "w" {
			numberAsString := arg[1:]
			weekOffset, err := strconv.Atoi(numberAsString)

			if err != nil {
                fmt.Printf("Could not parse %s as an integer\n", numberAsString)
                fmt.Println("Run 'glunch help' for help with available commands and their syntax")
                os.Exit(1)
			}

			cmd.GetMenu(weekOffset)
        } else if arg == "version" {
            printVersion()
		} else {
			printHelp()
		}
	} else {
		fmt.Println("Unknown number of arguments")
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Welcome to glunch - a lunch menu printer written in Go")
	fmt.Println("\nAvailable commands:")
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)
	fmt.Fprintln(w, "  glunch\tGets the menu for the current week highlighting today")
	fmt.Fprintln(w, "  glunch help\tShows this help menu")
	fmt.Fprintln(w, "  glunch version\tPrints the version of glunch")
	fmt.Fprintln(w, "  glunch w<int>\tGets the menu offset by the number of weeks for example 'glunch w-1'")
	fmt.Fprintln(w, "\tor 'glunch w1' to get the menu of the previous or next week respectively")
	w.Flush()
}

func printVersion() {
	fmt.Println("glunch v0.0.2")
}
