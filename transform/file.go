package transform

import (
    "github.com/ncipollo/fnew/message"
    "fmt"
    "errors"
    "io/ioutil"
)

type FileTransform struct {
    Options
    output   message.Printer
    globber  Globber
    replacer FileStringReplacer
}

func NewFileTransform(
    options Options,
    output message.Printer,
    globber Globber,
    replacer FileStringReplacer,
) *FileTransform {
    return &FileTransform{Options: options, output: output, globber: globber, replacer: replacer}
}

func (transform *FileTransform) Apply(variables Variables) error {
    inputPath := transform.Options.InputPath

    transform.output.Println(transform.transformMessage())

    paths, err := transform.globber.Glob(inputPath)
    if err != nil {
        return err
    }

    return transform.replaceStringsInPaths(paths, variables)
}

func (transform *FileTransform) transformMessage() string {
    inputPath := transform.Options.InputPath
    return fmt.Sprintf("Transforming file(s) (%s)", inputPath)
}

func (transform *FileTransform) replaceStringsInPaths(paths []string, variables Variables) error {
    stringReplace := transform.Options.StringReplace
    replacer := transform.replacer

    for _, filepath := range paths {
        err := replacer.Replace(filepath, stringReplace, variables)
        if err  != nil{
            return err
        }
    }

    return nil
}

type FileStringReplacer interface {
    Replace(filepath string, stringReplace StringReplace, variables Variables) error
}

type ioFileStringReplacer struct {}

func NewFileStringReplacer() FileStringReplacer {
    return &ioFileStringReplacer{}
}

func (*ioFileStringReplacer) Replace(filepath string, stringReplace StringReplace, variables Variables) error {
    data, err := ioutil.ReadFile(filepath)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to read file: %s", filepath))
    }

    fileString := string(data)
    replacedData := []byte(stringReplace.Replace(fileString, variables))

    err = ioutil.WriteFile(filepath, replacedData, 0777)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to write file: %s", filepath))
    } else {
        return nil
    }
}

