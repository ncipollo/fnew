package transform

import (
    "github.com/ncipollo/fnew/message"
    "os"
    "errors"
    "fmt"
)

type Factory struct {
    globber Globber
    output  message.Printer
}

func NewFactory() *Factory {
    return &Factory{
        globber: NewGlobber(),
        output:  message.NewStandardWriter(),
    }
}

func (factory *Factory) Transforms(optionsList []Options) ([]Transform, error) {
    var transforms []Transform

    for _, options := range optionsList {
        var err error
        var transform Transform
        switch options.Type {
        case TypeFileMove:
            transform = factory.createFileMoveTransform(options)
        case TypeFileStringReplace:
            transform = factory.createFileTransform(options)
        case TypeInput:
            transform = factory.createInputTransform(options)
        case TypeRunScript:
            transform = factory.createScriptTransform(options)
        case TypeVariableStringReplace:
            transform = factory.createVariableTransform(options)
        default:
            err = errors.New(fmt.Sprintf("invalid transform type: %s", options.Type))
        }

        if err != nil {
            return nil, err
        }

        transforms = append(transforms, transform)
    }

    return transforms, nil
}

func (factory *Factory) createFileMoveTransform(options Options) Transform {
    return NewFileMoveTransform(options, factory.output, NewFileMover())
}

func (factory *Factory) createFileTransform(options Options) Transform {
    return NewFileTransform(options, factory.output, factory.globber, NewFileStringReplacer())
}

func (factory *Factory) createInputTransform(options Options) Transform {
    return NewInputTransform(options, os.Stdin, factory.output)
}

func (factory *Factory) createScriptTransform(options Options) Transform {
    return NewScriptTransform(options, factory.output, NewCommandRunner())
}

func (factory *Factory) createVariableTransform(options Options) Transform {
    return NewVariableTransform(options, factory.output)
}
