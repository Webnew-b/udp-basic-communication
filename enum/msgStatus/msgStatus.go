package msgStatus

type MsgStatus uint16

const (
	SUCCESS          MsgStatus = 1
	UNKNOWN_ERROR    MsgStatus = 1000
	NOT_FOUND_CLIENT MsgStatus = 1001
)

func (status MsgStatus) String() string {
	switch status {
	case SUCCESS:
		return "Success"
	case UNKNOWN_ERROR:
		return "Unknown error"
	case NOT_FOUND_CLIENT:
		return "Not found client"
	default:
		return "Unknown"
	}
}
