package transform

type Transformer interface {
    Apply(transforms []Transform, variables Variables) error
}

type standardTransformer struct {}

func NewTransformer() Transformer {
    return &standardTransformer{}
}

func (*standardTransformer) Apply(transforms []Transform, variables Variables) error {
    for _, currentTransform := range transforms {
        err := currentTransform.Apply(variables)
        if err != nil {
            return err
        }
    }
    return nil
}
