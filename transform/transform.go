package transform

type Type string

const (
	TransformTypeFileMove              Type = "file_move"
	TransformTypeFileStringReplace     Type = "file_string_replace"
	TransformTypeInput                 Type = "input"
	TransformTypeRunScript             Type = "run_script"
	TransformTypeVariableStringReplace Type = "variable_string_replace"
)
