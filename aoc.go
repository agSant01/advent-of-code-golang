package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"
)

func init() {
	cli.HelpFlag = &cli.BoolFlag{Name: "help"}
	cli.VersionFlag = &cli.BoolFlag{Name: "version", Aliases: []string{"V"}}

	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Fprintf(cmd.Root().Writer, "version=%s\n", cmd.Root().Version)
	}
	cli.OsExiter = func(cmd int) {
		fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", cmd)
	}
	cli.FlagStringer = func(fl cli.Flag) string {
		return fmt.Sprintf("\t\t%s", fl.Names()[0])
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func create(day string, year string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/%s/day%s", cwd, year, day)
	println(path)

	log.Println("Creating day:", path)

	exists, err := pathExists(path)
	if err != nil {
		log.Fatalf("Error checking path: %s", err)
	}
	if exists {
		log.Fatalf("Exiting: File path for day %s/%s already exists...", year, day)
	}

	err = os.MkdirAll(path, os.FileMode(0755))
	println(os.FileMode(0755))
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(fmt.Sprintf("%s/main.go", path))
	if err != nil {
		log.Fatal(err)
	}
	mainDeclaration := fmt.Sprintf("package main;\n\nfunc main() {\n    println(\"This is main for %s/day%s\");\n}\n", year, day)
	file.WriteString(mainDeclaration)
	file.Close()

	_, err = os.Create(fmt.Sprintf("%s/data.txt", path))
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Create(fmt.Sprintf("%s/data_test.txt", path))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.SetPrefix("aoc-golang ")

	cmd := &cli.Command{
		Name:      "aoc-cli",
		Version:   "v0.0.1",
		Copyright: "MIT License",
		Usage:     "Golang Helper util to create and init days",
		UsageText: "contrive - demonstrating the available API",
		ArgsUsage: "[args and such]",
		Commands: []*cli.Command{
			{
				Name:      "create-day",
				Aliases:   []string{"c"},
				Usage:     "create a day",
				UsageText: "create-day [[--year|-y <YEAR>]] [[--day|-d <DAY>]]",
				// Description: "no really, there is a lot of dooing to be done",
				// ArgsUsage: "[arrgh]",
				// Flags: []cli.Flag{
				// 	&cli.StringFlag{Name: "day", Local: true, Aliases: []string{"d"}, Required: false, Value: time.Now().Format("02")},
				// 	&cli.StringFlag{Name: "year", Local: true, Aliases: []string{"y"}, Required: false, Value: time.Now().Format("2006")},
				// },
				ShellComplete: func(ctx context.Context, cmd *cli.Command) {
					fmt.Fprintf(cmd.Root().Writer, "--year|-y  --day|-d\n")
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					create(cmd.String("day"), cmd.String("year"))
					return nil
				},
				OnUsageError: func(ctx context.Context, cmd *cli.Command, err error, isSubcommand bool) error {
					fmt.Printf("%s\n", err)
					fmt.Fprintf(cmd.Root().Writer, "for shame\n")
					fmt.Println(cmd.String("year"))
					return err
				},
				EnableShellCompletion: true,
			},
			{
				Name:  "get-data",
				Usage: "downloads the data of the challenge",
				// UsageText:   "doo - does the dooing",
				// Description: "no really, there is a lot of dooing to be done",
				// ArgsUsage: "[arrgh]",

				ShellComplete: func(ctx context.Context, cmd *cli.Command) {
					fmt.Fprintf(cmd.Root().Writer, "--year|-y  --day|-d\n")
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					create(cmd.String("day"), cmd.String("year"))
					return nil
				},
				OnUsageError: func(ctx context.Context, cmd *cli.Command, err error, isSubcommand bool) error {
					fmt.Printf("%s\n", err)
					fmt.Fprintf(cmd.Root().Writer, "for shame\n")
					fmt.Println(cmd.String("year"))
					return err
				},
				EnableShellCompletion: true,
			},
		},
		Flags: []cli.Flag{
			// 	&cli.IntFlag{Name: "year", Aliases: []string{"y"}, Required: false, DefaultText: time.Now().Local().Format("YYYY")},
			// 	&cli.BoolFlag{Value: true, Name: "fancier"},
			// 	&cli.DurationFlag{Name: "howlong", Aliases: []string{"H"}, Value: time.Second * 3},
			// 	&cli.FloatFlag{Name: "howmuch"},
			// 	&cli.IntFlag{Name: "longdistance"},
			// 	&cli.IntSliceFlag{Name: "intervals"},
			// 	&cli.StringFlag{Name: "dance-move", Aliases: []string{"d"}},
			// 	&cli.StringSliceFlag{Name: "names", Aliases: []string{"N"}},
			// 	&cli.BoolFlag{Name: "run", Aliases: []string{"r"}},
			&cli.StringFlag{Name: "day", Local: true, Aliases: []string{"d"}, Required: false, Value: time.Now().Format("02")},
			&cli.StringFlag{Name: "year", Local: true, Aliases: []string{"y"}, Required: false, Value: time.Now().Format("2006")},
		},
		EnableShellCompletion: true,
		HideHelp:              false,
		HideVersion:           false,
		ShellComplete: func(ctx context.Context, cmd *cli.Command) {
			fmt.Fprintf(cmd.Root().Writer, "lipstick\nkiss\nme\nlipstick\nringo\n")
		},
		CommandNotFound: func(ctx context.Context, cmd *cli.Command, command string) {
			fmt.Fprintf(cmd.Root().Writer, "Thar be no %q here.\n", command)
		},
		OnUsageError: func(ctx context.Context, cmd *cli.Command, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}

			fmt.Fprintf(cmd.Root().Writer, "WRONG: %#v\n", err)
			return nil
		},
		// Action: func(ctx context.Context, cmd *cli.Command) error {
		// 	cli.DefaultAppComplete(ctx, cmd)
		// 	cli.HandleExitCoder(errors.New("not an exit coder, though"))
		// 	cli.ShowAppHelp(cmd)
		// 	cli.ShowCommandHelp(ctx, cmd, "also-nope")
		// 	cli.ShowSubcommandHelp(cmd)
		// 	cli.ShowVersion(cmd)

		// 	fmt.Printf("%#v\n", cmd.Root().Command("doo"))
		// 	if cmd.Bool("infinite") {
		// 		cmd.Root().Run(ctx, []string{"app", "doo", "wop"})
		// 	}

		// 	if cmd.Bool("forevar") {
		// 		cmd.Root().Run(ctx, nil)
		// 	}
		// 	fmt.Printf("%#v\n", cmd.Root().VisibleCategories())
		// 	fmt.Printf("%#v\n", cmd.Root().VisibleCommands())
		// 	fmt.Printf("%#v\n", cmd.Root().VisibleFlags())

		// 	fmt.Printf("%#v\n", cmd.Args().First())
		// 	if cmd.Args().Len() > 0 {
		// 		fmt.Printf("%#v\n", cmd.Args().Get(1))
		// 	}
		// 	fmt.Printf("%#v\n", cmd.Args().Present())
		// 	fmt.Printf("%#v\n", cmd.Args().Tail())

		// 	ec := cli.Exit("ohwell", 86)
		// 	fmt.Fprintf(cmd.Root().Writer, "%d", ec.ExitCode())
		// 	fmt.Printf("made it!\n")
		// 	return ec
		// },
		Metadata: map[string]interface{}{
			"layers":          "many",
			"explicable":      false,
			"whatever-values": 19.99,
		},
	}

	cmd.Run(context.Background(), os.Args)
}
