package gotchas

var (
	_ error = (*CustomPointerReceiverError)(nil)
	_ error = CustomValueReceiverError{}
	_ error = (*CustomValueReceiverError)(nil)
)

type CustomPointerReceiverError struct {
}

func (p *CustomPointerReceiverError) Error() string {
	return "custom error implementing by pointer receiver"
}

type CustomValueReceiverError struct {
}

func (v CustomValueReceiverError) Error() string {
	return "custom error implementing by value receiver"
}
