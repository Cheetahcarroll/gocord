package models

// ApplicationFlag is the int alias for application flags
type ApplicationFlag int

const (
	// GatewayPresence is the GATEWAY_PRESENCE flag
	GatewayPresence ApplicationFlag = 1 << (iota + 12)
	// GatewayPresenceLimited is the GATEWAY_PRESENCE_LIMITED flag
	GatewayPresenceLimited
	// GatewayGuildMembers is the GATEWAY_GUILD_MEMBERS flag
	GatewayGuildMembers
	// GatewayGuildMembersLimited is the GATEWAY_GUILD_MEMBERS_LIMITED flag
	GatewayGuildMembersLimited
	// VerificationPendingGuildLimit is the VERIFICATION_PENDING_GUILD_LIMIT flag
	VerificationPendingGuildLimit
	// Embedded is the EMBEDDED flag
	Embedded
)

// Application models the discord application object
type Application struct {
	ID                  Snowflake       `json:"id"`
	Name                string          `json:"name"`
	Icon                string          `json:"icon"`
	Description         string          `json:"description"`
	RPCOrigins          []string        `json:"rpc_origins,omitempty"`
	BotPublic           bool            `json:"bot_public"`
	BotRequireCodeGrant bool            `json:"bot_require_code_grant"`
	TermsOfServiceURL   string          `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL    string          `json:"privacy_policy_url,omitempty"`
	Owner               User            `json:"owner,omitempty"`
	Summary             string          `json:"summary"`
	VerifyKey           string          `json:"verify_key"`
	Team                *Team           `json:"team"`
	GuildID             Snowflake       `json:"guild_id,omitempty"`
	PrimarySkuID        Snowflake       `json:"primary_sku_id,omitempty"`
	Slug                string          `json:"slug,omitempty"`
	CoverImage          string          `json:"cover_image,omitempty"`
	Flags               ApplicationFlag `json:"flags,omitempty"`
}
