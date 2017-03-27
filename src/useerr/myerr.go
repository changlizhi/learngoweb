package useerr

type SyntaxError struct {
	msg    string
	Offset int64
}

func (e *SyntaxError) Error() string {
	return e.msg
}

func Use_my_err() {
	//if err := Decode
}
