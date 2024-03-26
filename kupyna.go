package kupyna

import (
    "fmt"
)

type KeySizeError int

func (k KeySizeError) Error() string {
    return fmt.Sprintf("kupyna256: invalid key size %d", int(k))
}
