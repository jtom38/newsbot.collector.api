package output

type Output interface {
	GeneratePayload() error
	SendPayload() error
}
