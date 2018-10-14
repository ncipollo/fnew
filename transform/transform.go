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
	Arguments []string `json:"arguments,omitempty"`
	InputPath string `json:"input_path,omitempty"`
	InputVariable string `json:"input_variable,omitempty"`
	OutputPath string `json:"output_path,omitempty"`
	OutputVariable string `json:"output_variable,omitempty"`
	StringReplace StringReplace `json:"string_replace,omitempty"`
	Type Type `json:"type"`
}

type StringReplace struct {
	Old string `json:"old"`
	New string `json:"new"`
}
