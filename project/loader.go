package project

type Loader interface {
    Load(filename string) (*Project, error)
}

type fileLoader struct { }

func NewLoader() Loader {
    return fileLoader{}
}

func (fileLoader) Load(filename string) (*Project, error) {
    return FromFile(filename)
}


