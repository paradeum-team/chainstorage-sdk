package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cmdDescription string = `
USAGE
  gcscmd  - Golang ChainStorage Command line tool

  gcscmd [--config=<config> | -c] [--debug | -D] [--help] [-h]  [--timeout=<timeout>] <command> ...

SUBCOMMANDS
  BASIC COMMANDS
    ls <ref>        List links from object or bucket
    mb  <ref>       Create bucket
    rb <ref>        Delete bucket
    get <ref>       Download objects
    put  <ref>      Upload objects 
    rm  <ref>       Delete objects or Clear bucket
    rn  <ref>       Rename objects name
    Import <ref>    Import carfile

  TOOL COMMANDS
    config        Manage configuration
    version       Show IPFS version information
    log           Manage and show logs of running daemon

  Use 'gcscmd <command> --help' to learn more about each command.

  ggcscmd uses the local file system. By default, Path is
  located at ~/. gcscmd. To change config, set the $GCSCMD_PATH
  environment variable:

    export GCSCMD_PATH =/path/to/gcscmd

  EXIT STATUS

  The CLI will exit with one of the following values:

  0     Successful execution.
  1     Failed executions.`

var rootCmd = &cobra.Command{
	Use:   "gcscmd",
	Short: "gcscmd",
	//Long:  cmdDescription,
	//Example: `gcscmd [--config=<config> | -c] [--debug | -D] [--help] [-h]  [--timeout=<timeout>] <command> ...`,
	Long: `gcscmd - Golang ChainStorage Command line tool`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {

	rootCmd.AddCommand(
		bucketListCmd,
		bucketCreateCmd,
		bucketRemoveCmd,
		bucketEmptyCmd,
		objectListCmd,
		objectRenameCmd,
		objectRemoveCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
