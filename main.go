package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	config     *DNSConfig
	configPath = "./blocking.json"
)

// Custom usage message
const usgMsg = "Commands:\n" +
	"\t Running without arguments Updates and Merges the blocklists.\n" +
	"\t update Updates and Merges the blocklists.\n" +
	"\t add -c <category> <blocklists> Adds the following blocklists to the config.\n" +
	"\t remove -c <category> <blocklists> Removes the blocklists.\n" +
	"The blocklists must be given with their full Path or URL!\n"

func main() {
	// load the config first thing!
	config = LoadConfig("./blocking.json")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCategory := addCmd.String("c", "", "Choose a category: ads, adult, etc...")

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	remCategory := removeCmd.String("c", "", "Choose a category: ads, adult, etc...")

	// enableCmd := flag.NewFlagSet("enable", flag.ExitOnError)

	// disableCmd := flag.NewFlagSet("disable", flag.ExitOnError)

	flag.Usage = func() {
		fmt.Printf("Usage: %s [command] [args]\n", os.Args[0])
		fmt.Print(usgMsg)
	}

	flag.Parse()

	// just running the binary updates and merges the blocklists no questions asked
	if len(os.Args) < 2 {
		fmt.Printf("No arguments, default behaviour is to update all blocklists!\n\n")
		UpdateListsAndMergeTags(config, "./dns")
		os.Exit(0)
	}

	// TODO: Need to also check for valid category and for empty arguments!
	switch os.Args[1] {
	case "update":
		updateCmd.Parse(os.Args[2:])
		fmt.Println("Updating the blocklists...")
		UpdateListsAndMergeTags(config, "./dns")
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Printf("Adding new blocklist in category %s\n", *addCategory)
		fmt.Println(addCmd.Args())
		fmt.Println("NOT IMPLEMENTED YET!")
		// add function
	case "remove":
		removeCmd.Parse(os.Args[2:])
		fmt.Printf("Removing blocklist from category %s\n", *remCategory)
		fmt.Println(removeCmd.Args())
		fmt.Println("NOT IMPLEMENTED YET!")
		// remove function
	default:
		flag.Usage()
		os.Exit(1)
	}
}
