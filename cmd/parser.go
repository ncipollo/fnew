package cmd

import (
    "flag"
    "fmt"
    "github.com/ncipollo/fnew/action"
    "path/filepath"
)

const listUsage = "lists the available fnew projects"
const updateUsage = "updates the fnew project list"
const shortHandSuffix = " (shorthand)"

type Parser struct {
    env []string
}

func NewParser(env []string) *Parser {
    return &Parser{env: env}
}

func (parser *Parser) Parse() Command {
    list := parser.setupListFlag()
    update := parser.setupUpdateFlag()

    flag.Parse()

    if *list {
        fmt.Println("List: ", *list)
        actionFactory := action.NewFactory("", "")
        return parser.listCommand(actionFactory)
    } else if *update {
        fmt.Println("Update: ", *update)
        actionFactory := action.NewFactory("", "")
        return parser.updateCommand(actionFactory)
    } else {
        // TODO: Check arg count (2)
        localPath, _ := parser.localPath()
        projectName := parser.projectName()
        actionFactory := action.NewFactory(localPath, projectName)

        fmt.Println("Local Path: ", localPath)
        fmt.Println("Project Name: ", projectName)

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

func (Parser) localPath() (string, error) {
    relativePath := flag.Arg(1)
    return filepath.Abs(relativePath)
}

func (Parser) projectName() string {
    return flag.Arg(0)
}

func (Parser) setupListFlag() *bool {
    var list bool
    flag.BoolVar(&list, "list", false, listUsage)
    flag.BoolVar(&list, "l", false, listUsage+shortHandSuffix)

    return &list
}

func (Parser) setupUpdateFlag() *bool {
    var list bool
    flag.BoolVar(&list, "update", false, updateUsage)
    flag.BoolVar(&list, "u", false, updateUsage+shortHandSuffix)

    return &list
}
