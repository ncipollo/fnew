package action

import (
    "github.com/ncipollo/fnew/message"
)

type Action interface {
    Perform(output message.Printer) error
}
