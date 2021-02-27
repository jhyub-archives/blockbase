package main

import (
	"fmt"
	"github.com/blbase/serve"
	"github.com/jessevdk/go-flags"
	"os"
)

type runOpts struct {
	Help   bool   `short:"h" long:"help"`
	Daemon bool   `short:"d" long:"daemon"`
	Port   int    `long:"port" default:"8080"`
	Host   string `long:"host" default:"0.0.0.0"`
}

type rootOpts struct {
	Help    bool    `short:"h" long:"help"`
	Version bool    `long:"version"`
	Run     runOpts `command:"run"`
}

var opts = new(rootOpts)

var version string = "1.0.0"

func main() {
	p := flags.NewParser(opts, flags.None)

	_, err := p.Parse()
	if p.Active != nil {
		if err != nil {
			PrintHelp(p.Active.Name)
			os.Exit(0)
		}
		switch p.Active.Name {
		case "run":
			RunCommand()
		}
	} else {
		RootCommand()
	}
}

func RunCommand() {
	if opts.Run.Help {
		print(3)
		PrintHelp("run")
		os.Exit(0)
	} else if opts.Run.Daemon {
		print(2)
		fmt.Println("Daemon not supported yet")
		os.Exit(0)
	} else {
		serve.Serve(fmt.Sprintf("%s:%d", opts.Run.Host, opts.Run.Port))
		os.Exit(0)
	}
}

func RootCommand() {
	if opts.Help {
		PrintHelp("root")
		os.Exit(0)
	} else if opts.Version {
		fmt.Println("Blockbase version v" + version)
		fmt.Println("Created by timtermtube@github.com. Maintained by organization blbase@github.com.")
		fmt.Println("Distributed as open-source under the MIT License.")
		fmt.Println("View contributors: https://github.com/blbase/blockbase/graphs/contributors")
		os.Exit(0)
	} else {
		PrintHelp("root")
		os.Exit(0)
	}
}

func PrintHelp(subcommand string) {
	switch subcommand {
	case "root":
		fmt.Println("Blockbase - New Model of Database")
		fmt.Println("")
		fmt.Println("Usage: blockbase [options] [COMMAND]")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -h, --help HELP         Display this message.")
		fmt.Println("  --version VERSION       Display version, license, authors.")
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  run        Start up blockbase server.")
	case "run":
		fmt.Println("Starts up the blockbase server.")
		fmt.Println("")
		fmt.Println("Usage: blockbase run [options]")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -h, --help HELP         Display this message.")
		fmt.Println("  -d, --daemon DAEMON     Run blockbase as daemon.")
		fmt.Println("  --port PORT             Set port of blockbase. Default value is 8080")
		fmt.Println("  --host HOST             Set hostname of blockbase. Default value is 0.0.0.0")
	}
}
