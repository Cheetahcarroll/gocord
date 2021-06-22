package models

// MembershipState is the int alias for Membership States
type MembershipState int

const (
	// Invited indicates the team member has been invited
	Invited = MembershipState(1)
	// Accepted indicates the team member accepted the invitation
	Accepted = MembershipState(2)
)

// Team models the discord team object
type Team struct {
	ID          Snowflake    `json:"id"`
	Icon        *string      `json:"icon"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID Snowflake    `json:"owner_user_id"`
}

// TeamMember models the discord team member object
type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"`
	Permissions     []string        `json:"permissions"`
	TeamID          Snowflake       `json:"team_id"`
	User            User            `json:"user"`
}
