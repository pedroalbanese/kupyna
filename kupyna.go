package kupyna

import (
    "fmt"
)

type KeySizeError int

func (k KeySizeError) Error() string {
    return fmt.Sprintf("kupyna: invalid key size %d", int(k))
}
