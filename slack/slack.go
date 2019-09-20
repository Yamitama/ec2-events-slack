package slack

type ec2EventSlackMessage struct {
	payload ec2EventPayload
	url     string
}

type ec2EventPayload struct {
}

// NewEC2EventSlackMessage returns a ec2EventSlackMessage struct with specified parameters
func NewEC2EventSlackMessage(
	channel string,
	username string,
	iconEmoji string,
	pretext string,
	title string,
	titleLink string,
	text string,
	resourceName string,
	resourceID string,
	eventType string,
	eventStatus string,
	eventDescription string,
	startTime string) {

}
