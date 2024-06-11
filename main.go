//go:build windows

/*
Copyright © 2024 Ryan Gravlin ryan.gravlin@gmail.com
*/
package main

import (
	"github.com/rgravlin/noitabackup/pkg/cmd"
)

func main() {
	cmd.Execute()
}

// TODO: Implement the following features:
//  * Viper configuration for: Steam path
//  * Feedback for UI async backup/restore: click -> animated processing -> success/fail
