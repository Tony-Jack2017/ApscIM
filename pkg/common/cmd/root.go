package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type ApscCommand struct {
	Command *cobra.Command
	Name    string
}

func init() {
	apscCommand := &ApscCommand{}
	apscCommand.Name = "APSC"
	apscCommand.Command = &cobra.Command{
		Use:   "apsc",
		Short: "APSC is IM system for private",
		Long: `APSC is a open source im system, private can use it build own
internal chat system for free. like other chat system which have some interesting
function, enjoy yourself
			`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Cobra !!!")
		},
	}

	start := NewStartCommand()
	version := NewVersionCommand()
	check := NewCheckCommand()

	apscCommand.Command.AddCommand(start.Command)
	apscCommand.Command.AddCommand(version.Command)
	apscCommand.Command.AddCommand(check.Command)
}

func NewStartCommand() *ApscCommand {
	startCommand := &ApscCommand{}
	startCommand.Name = "start"
	startCommand.Command = &cobra.Command{
		Use:   "start",
		Short: "start subcommand for staring the im system",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Cobra !!!")
		},
	}
	return startCommand
}

func NewVersionCommand() *ApscCommand {
	versionCommand := &ApscCommand{}
	versionCommand.Name = "version"
	versionCommand.Command = &cobra.Command{
		Use:   versionCommand.Name,
		Short: "version subcommand show apsc version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("apsc v1.0.0")
		},
	}
	return versionCommand
}

func NewCheckCommand() *ApscCommand {
	checkCommand := &ApscCommand{}
	checkCommand.Name = "check"
	checkCommand.Command = &cobra.Command{
		Use:   checkCommand.Name,
		Short: "check subcommand is for check every serve status",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("apsc v1.0.0")
		},
	}
	return checkCommand
}

func NewSubCommand(name string, des string, run func(cmd *cobra.Command, args []string)) {
	cmd := &ApscCommand{}
	cmd.Name = name
	cmd.Command = &cobra.Command{
		Use:   cmd.Name,
		Short: des,
		Run:   run,
	}
}

func AddCommands(cmds ...*cobra.Command) {
	root := &ApscCommand{}
	root.Command.AddCommand(cmds...)
}
