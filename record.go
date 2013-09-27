package main

import (
)

var cmdRecord  = &Command{
	Run:   runRecord,
	Usage: "record <command> [<args>]",
	Short: "Create, modify, or view records",
	Long: `
Create, modify, or view records

Usage:

  force record get <object> <id>

  force record create <object> [<fields>]

  force record update <object> <id> [<fields>]

Examples:

  force record get User 00Ei0000000000

  force record create User Name:"David Dollar" Phone:0000000000

  force record update User 00Ei0000000000 State:GA
`,
}

func runRecord(cmd *Command, args []string) {
	switch args[0] {
	case "get":
		runRecordGet(args[1:])
	case "create":
		runRecordCreate(args[1:])
	case "update":
		runRecordUpdate(args[1:])
	default:
		ErrorAndExit("so such subcommand for record: %s", args[0])
	}
}

func runRecordGet(args []string) {
	force, _ := ActiveForce()
	if len(args) != 2 {
		ErrorAndExit("must specify object and id")
	}
	object, err := force.GetRecord(args[0], args[1])
	if err != nil {
		ErrorAndExit(err.Error())
	} else {
		DisplayForceRecord(object)
	}
}

func runRecordCreate(args []string) {
	if len(args) < 1 {
		ErrorAndExit("must specify object")
	}
	ErrorAndExit("not implemented yet")
}

func runRecordUpdate(args []string) {
	ErrorAndExit("not implemented yet")
}
