package events

type AMMEvent string

const (
	REFRESH_CONFIG AMMEvent = "refresh_config"
)

var AMMEvents = []struct {
	Value  AMMEvent
	TSName string
}{
	{REFRESH_CONFIG, "REFRESH_CONFIG"},
}
