package utility

import (
	"fmt"
)

//Error handling
type Myerr struct {
	Origin  string
	Message error
}

func (me *Myerr) Error() string {
	return fmt.Sprintf("api.%s()|%s", me.Origin, me.Message)
}
