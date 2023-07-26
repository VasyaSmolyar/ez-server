package action

type NotFoundError struct{}

func (err *NotFoundError) Error() string {
	return "Record not found"
}
