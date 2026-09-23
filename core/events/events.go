package events

type AMMEvent string

const (
	REFRESH_CONFIG AMMEvent = "refresh-config"
)

var AMMEvents = []struct {
	Value  AMMEvent
	TSName string
}{
	{REFRESH_CONFIG, "refresh-config"},
}
