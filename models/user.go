package models

import "strings"

// UserFlag is an int alias for user flags
type UserFlag int

// VisibilityType is an int alias for visibility
type VisibilityType int

const (
	// Indicates a user is a discord employee
	Employee UserFlag = 1 << iota
	// Indicates a user is partnered server owner
	Partner
	// Indicates a user is on the hype squad
	HypeSquad
	// Indicates a user is a level 1 bug hunter
	BugHunter1
	_
	_
	// Indicates a user is part of house bravery
	Bravery
	// Indicates a user is part of house brilliance
	Brilliance
	// Indicates a user is part of house balance
	Balance
	// Indicates a user is an early supporter
	EarlySupporter
	// Indicates a user is a team user
	TeamUser
	_
	_
	_
	// Indicates a user is a level 2 bug hunter
	BugHunter2
	_
	// Indicates a user is a verified bot
	VerifiedBot
	// Indicates a user is an early verified bot developer
	EarlyVerifiedBotDeveloper
	// Indicates a user is a discord certified moderator
	DiscordCertifiedModerator

	// Invisible indicates the message is only visible to the user
	Invisible = VisibilityType(0)
	// Visible indicates the message is visible to everyone
	Visible = VisibilityType(1)
)

// User is a struct that models the discord user object
type User struct {
	ID            Snowflake `json:"id"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Avatar        *string   `json:"avatar"`
	Bot           bool      `json:"bot,omitempty"`
	System        bool      `json:"system,omitempty"`
	MFA           bool      `json:"mfa_enabled,omitempty"`
	Locale        string    `json:"locale,omitempty"`
	Verified      bool      `json:"verified,omitempty"`
	Email         *string   `json:"email,omitempty"`
	Flags         UserFlag  `json:"flags,omitempty"`
	PremiumType   int       `json:"premium_type,omitempty"`
}

// Connection models the discord connection object
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

// String returns the represented flag value
func (f UserFlag) String() string {
	if f&Employee == Employee {
		return "Discord Employee"
	}
	if f&Partner == Partner {
		return "Discord Partner"
	}
	if f&HypeSquad == HypeSquad {
		return "Hype Squad Events"
	}
	if f&BugHunter1 == BugHunter1 {
		return "Bug Hunter Level 1"
	}
	if f&Bravery == Bravery {
		return "House Bravery"
	}
	if f&Brilliance == Brilliance {
		return "House Brilliance"
	}
	if f&Balance == Balance {
		return "House Balance"
	}
	if f&EarlySupporter == EarlySupporter {
		return "Early Supporter"
	}
	if f&TeamUser == TeamUser {
		return "Team User"
	}
	if f&BugHunter2 == BugHunter2 {
		return "Bug Hunter Level 2"
	}
	if f&VerifiedBot == VerifiedBot {
		return "Verified Bot"
	}
	if f&EarlyVerifiedBotDeveloper == EarlyVerifiedBotDeveloper {
		return "Early Verified Bot Developer"
	}
	if f&DiscordCertifiedModerator == DiscordCertifiedModerator {
		return "Discord Certified Moderator"
	}
	return "None"
}

// ValidateUsername ensures a username is valid
func ValidateUsername(username string) bool {
	username = strings.ReplaceAll(strings.Trim(username, " "), "  ", " ")

	if len(username) > 32 || len(username) < 2 {
		return false
	}

	if strings.ContainsAny(username, "@#:") || strings.Contains(username, "```") {
		return false
	}

	switch username {
	case "discordtag":
		return false
	case "everyone":
		return false
	case "here":
		return false
	}

	return true
}

// ValidateNickname ensures a nickname is valid
func ValidateNickname(nickname string) bool {
	nickname = strings.ReplaceAll(strings.Trim(nickname, " "), "  ", " ")

	if len(nickname) > 32 || len(nickname) == 0 {
		return false
	}

	if strings.ContainsAny(nickname, "@#:") || strings.Contains(nickname, "```") {
		return false
	}

	return true
}
