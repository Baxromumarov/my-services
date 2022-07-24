package messagebroker

type Consumer interface {
	Start() error
}