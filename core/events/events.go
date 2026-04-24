package events

type AMMEvent string

const (
	REFRESH_CONFIG AMMEvent = "config:refresh"
)

var AMMEvents = []struct {
	Value  AMMEvent
	TSName string
}{
	{REFRESH_CONFIG, "config:refresh"},
}
