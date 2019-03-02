package transform

import (
    "strings"
    "os"
)

type Type string

const (
    TypeFileMove              Type = "file_move"
    TypeFileStringReplace     Type = "file_string_replace"
    TypeInput                 Type = "input"
    TypeRunScript             Type = "run_script"
    TypeVariableStringReplace Type = "variable_string_replace"
)

type Transform interface {
    Apply(variables Variables) error
}

type Options struct {
    Arguments            []string      `json:"arguments,omitempty"`
    InputPath            string        `json:"input_path,omitempty"`
    InputVariable        string        `json:"input_variable,omitempty"`
    OutputPath           string        `json:"output_path,omitempty"`
    OutputVariable       string        `json:"output_variable,omitempty"`
    SkipIfVariableExists bool          `json:"skip_if_variable_exists,omitempty"`
    StringPrefix         string        `json:"string_prefix,omitempty"`
    StringReplace        StringReplace `json:"string_replace,omitempty"`
    StringSuffix         string        `json:"string_suffix,omitempty"`
    Type                 Type          `json:"type"`
}

type StringReplace struct {
    Old string `json:"old"`
    New string `json:"new"`
}

func (stringReplace *StringReplace) Replace(input string, variables Variables) string {
    return strings.Replace(input,
        stringReplace.oldString(variables),
        stringReplace.newString(variables),
        -1)
}

func (stringReplace *StringReplace) oldString(variables Variables) string {
    if strings.HasPrefix(stringReplace.Old, "$") {
        return variables[strings.TrimPrefix(stringReplace.Old, "$")]
    } else {
        return stringReplace.Old
    }
}

func (stringReplace *StringReplace) newString(variables Variables) string {
    if strings.HasPrefix(stringReplace.New, "$") {
        return variables[strings.TrimPrefix(stringReplace.New, "$")]
    } else {
        return stringReplace.New
    }
}

type Variables map[string]string

func NewVariables() Variables {
    return make(Variables)
}

func (variables Variables) AddEnv() {
    for _, env := range os.Environ() {
        parts := strings.Split(env, "=")
        key := parts[0]
        value := parts[1]
        variables[key] = value
    }
}

func (variables Variables) AddProjectName(projectName string) {
    variables["project_name"] = projectName
}
