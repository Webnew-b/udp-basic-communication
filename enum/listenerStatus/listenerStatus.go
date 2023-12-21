package listenerStatus

type Status uint8

const (
	STOP       Status = 0
	START      Status = 1
	READY_STOP Status = 2
)

func (status Status) String() string {
	switch status {
	case STOP:
		return "This listener is stopped"
	case START:
		return "This listener is starting"
	case READY_STOP:
		return "This listener is reading to stop"
	default:
		return "Unknown status"
	}
}
