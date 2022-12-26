package stub

import "errors"

func UnimplementedMethodCalled() error {
	return errors.New("error: unimplemented method called.")
}
