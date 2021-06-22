package models

import "time"

// ExpireBehavior is an int alias for expiration behavior
type ExpireBehavior int

const (
	// RemoveRole is the behavior for removing a role from a user or bot
	RemoveRole = ExpireBehavior(0)
	// Kick is the behavior for kicking a user or bot
	Kick = ExpireBehavior(1)
)

// Integration models the integration information for discord
type Integration struct {
	ID                Snowflake              `json:"id"`
	Name              string                 `json:"name"`
	Type              string                 `json:"type"`
	Enabled           bool                   `json:"enabled"`
	Syncing           bool                   `json:"syncing,omitempty"`
	RoleID            Snowflake              `json:"role_id,omitempty"`
	EnableEmoticons   bool                   `json:"enable_emoticons,omitempty"`
	ExpireBehavior    ExpireBehavior         `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int                    `json:"expire_grace_period,omitempty"`
	User              User                   `json:"user,omitempty"`
	Account           IntegrationAccount     `json:"account"`
	SyncedAt          time.Time              `json:"synced_at,omitempty"`
	SubscriberCount   int                    `json:"subscriber_count,omitempty"`
	Revoked           bool                   `json:"revoked,omitempty"`
	Application       IntegrationApplication `json:"application,omitempty"`
}

// IntegrationAccount models the account use with Integration
type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// IntegrationApplication models the application used for integration
type IntegrationApplication struct {
	ID          Snowflake `json:"id"`
	Name        string    `json:"name"`
	Icon        *string   `json:"icon"`
	Description string    `json:"description"`
	Summary     string    `json:"summary"`
	User        User      `json:"bot,omitempty"`
}
