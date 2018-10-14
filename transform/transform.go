package transform

type Type string

const (
	TransformTypeFileMove              Type = "file_move"
	TransformTypeFileStringReplace     Type = "file_string_replace"
	TransformTypeInput                 Type = "input"
	TransformTypeRunScript             Type = "run_script"
	TransformTypeVariableStringReplace Type = "variable_string_replace"
)

type Transform interface {
	Apply()
}

type TransformOptions struct {
	Arguments []string `json:"arguments"`
	InputPath *string `json:"input_path"`
	InputVariable *string `json:"input_variable"`
	OutputPath *string `json:"output_path"`
	OutputVariable *string `json:"output_variable"`
	StringReplace *StringReplace `json:"string_replace"`
}

type StringReplace struct {
	Old string
	New string
}
