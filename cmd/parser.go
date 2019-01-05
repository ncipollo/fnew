package cmd

import (
    "flag"
    "fmt"
    "github.com/ncipollo/fnew/action"
    "path/filepath"
    "os"
)

const listUsage = "lists the available fnew projects"
const updateUsage = "updates the fnew project list"
const shortHandSuffix = " (shorthand)"

type Parser struct {
    env  []string
    flag *flag.FlagSet
}

func NewParser(env []string) *Parser {
    name := os.Args[0]
    flagSet := flag.NewFlagSet(name, flag.ExitOnError)
    flagSet.Usage = func() {
        fmt.Fprintf(flagSet.Output(),
            "Usage: %s [OPTIONS] <source project> <project location>\n",
            name)
        flagSet.PrintDefaults()
    }
    return &Parser{env: env, flag: flagSet}
}

func (parser *Parser) Parse() Command {
    list := parser.setupListFlag()
    update := parser.setupUpdateFlag()

    parser.flag.Parse(os.Args[1:])

    if *list {
        fmt.Println("List: ", *list)
        actionFactory := action.NewFactory("", "")
        return parser.listCommand(actionFactory)
    } else if *update {
        fmt.Println("Update: ", *update)
        actionFactory := action.NewFactory("", "")
        return parser.updateCommand(actionFactory)
    } else {
        if parser.flag.NArg() < 2 {
            parser.printUsageAndExit()
        }
        localPath, _ := parser.localPath()
        projectName := parser.projectName()
        actionFactory := action.NewFactory(localPath, projectName)

        fmt.Println("Project Name: ", projectName)
        fmt.Println("Local Path: ", localPath)

        return parser.createCommand(actionFactory)
    }

}

func (Parser) createCommand(actionFactory *action.Factory) Command {
    return nil
}

func (Parser) listCommand(actionFactory *action.Factory) Command {
    return nil
}

func (Parser) updateCommand(actionFactory *action.Factory) Command {
    return nil
}

func (parser *Parser) localPath() (string, error) {
    relativePath := parser.flag.Arg(1)
    return filepath.Abs(relativePath)
}

func (parser *Parser) projectName() string {
    return parser.flag.Arg(0)
}

func (parser *Parser) printUsageAndExit() {
    parser.flag.Usage()
    os.Exit(2)
}

func (parser *Parser) setupListFlag() *bool {
    var list bool
    parser.flag.BoolVar(&list, "list", false, listUsage)
    parser.flag.BoolVar(&list, "l", false, listUsage+shortHandSuffix)

    return &list
}

func (parser *Parser) setupUpdateFlag() *bool {
    var list bool
    parser.flag.BoolVar(&list, "update", false, updateUsage)
    parser.flag.BoolVar(&list, "u", false, updateUsage+shortHandSuffix)

    return &list
}
