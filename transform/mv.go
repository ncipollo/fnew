package transform

import (
    "github.com/ncipollo/fnew/message"
    "fmt"
    "errors"
    "os"
    "strings"
    "path/filepath"
)

type FileMoveTransform struct {
    Options
    output message.Printer
    mover  FileMover
}

func NewFileMoveTransform(options Options, output message.Printer, mover FileMover) *FileMoveTransform {
    return &FileMoveTransform{Options: options, output: output, mover: mover}
}

func (transform *FileMoveTransform) Apply(variables Variables) error {
    transform.output.Println(transform.transformMessage(variables))

    err := transform.validateInputPath(variables)
    if err != nil {
        return err
    }

    err = transform.validateOutputPathPath(variables)
    if err != nil {
        return err
    }

    return transform.mover.Move(transform.inputPath(variables), transform.outputPath(variables))
}

func (transform *FileMoveTransform) transformMessage(variables Variables) string {
    inputPath := transform.inputPath(variables)
    outputPath := transform.outputPath(variables)

    return fmt.Sprintf("Moving file(s) from (%s), to (%s)", inputPath, outputPath)
}

func (transform *FileMoveTransform) validateInputPath(variables Variables) error {
    inputPath := transform.inputPath(variables)
    if len(inputPath) == 0 {
        return errors.New("invalid input path")
    }
    return nil
}

func (transform *FileMoveTransform) validateOutputPathPath(variables Variables) error {
    inputPath := transform.outputPath(variables)
    if len(inputPath) == 0 {
        return errors.New("invalid output path")
    }
    return nil
}

func (transform *FileMoveTransform) inputPath(variables Variables) string {
    if strings.HasPrefix(transform.Options.InputPath, "$") {
        return variables[strings.TrimPrefix(transform.Options.InputPath, "$")]
    } else {
        return transform.Options.InputPath
    }
}

func (transform *FileMoveTransform) outputPath(variables Variables) string {
    if strings.HasPrefix(transform.Options.OutputPath, "$") {
        return variables[strings.TrimPrefix(transform.Options.OutputPath, "$")]
    } else {
        return transform.Options.OutputPath
    }
}

type FileMover interface {
    Move(oldPath string, newPath string) error
}

type osFileMover struct{}

func NewFileMover() FileMover {
    return &osFileMover{}
}

func (mover *osFileMover) Move(oldPath string, newPath string) error {
    absOldPath, err := filepath.Abs(oldPath)
    if err != nil {
        return err
    }
    absNewPath, err := filepath.Abs(newPath)
    if err != nil {
        return err
    }
    absNewDir, err := filepath.Abs(filepath.Join(absNewPath, ".."))
    if err != nil {
        return err
    }
    err = os.MkdirAll(absNewDir, 0777)
    if err != nil {
        return err
    }
    err = os.Rename(absOldPath, absNewPath)
    if err != nil {
        return err
    }
    return mover.removeEmptyDirectory(absOldPath)
}

func (osFileMover) removeEmptyDirectory(absOldPath string) error {
    absOldDir, err := filepath.Abs(filepath.Join(absOldPath, ".."))
    if err != nil {
        return err
    }

    var notEmpty bool
    err = filepath.Walk(absOldDir, func(path string, info os.FileInfo, err error) error {
        if err == nil && path != absOldDir {
            notEmpty = true
        }
        return err
    })

    if !notEmpty {
        return os.RemoveAll(absOldDir)
    }

    return nil
}
