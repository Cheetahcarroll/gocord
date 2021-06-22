package models

// VisibilityType is an int alias for visibility
type VisibilityType int

const (
	// Invisible indicates the message is only visible to the user
	Invisible = VisibilityType(0)
	// Visible indicates the message is visible to everyone
	Visible = VisibilityType(1)
)

// Connection is for user connections
type Connection struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Revoked      bool           `json:"revoked,omitempty"`
	Integrations []string       `json:"integrations,omitempty"`
	Verified     bool           `json:"verified"`
	FriendSync   bool           `json:"friend_sync"`
	ShowActivity bool           `json:"show_activity"`
	Visibility   VisibilityType `json:"visibility"`
}
