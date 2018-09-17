package action

import "io"

type Action interface {
	Perform(output io.Writer) error
}
