package transform

import "strings"

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
    StringReplace        StringReplace `json:"string_replace,omitempty"`
    Type                 Type          `json:"type"`
}

type StringReplace struct {
    Old string `json:"old"`
    New string `json:"new"`
}

func (stringReplace *StringReplace) Replace(input string) string {
    return strings.Replace(input, stringReplace.Old, stringReplace.New, -1)
}

type Variables map[string]string

func NewVariables() Variables {
    return make(Variables)
}
