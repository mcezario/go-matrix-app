package parsers

type ParserError struct {
	Err error
}

func (err *ParserError) Error() string {
	return err.Err.Error()
}
