package array_utils

type ArrayUtilsError struct {
	Err error
}

func (err *ArrayUtilsError) Error() string {
	return err.Err.Error()
}
