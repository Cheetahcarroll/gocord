package models

type AuditLogEvent int

const (
	// GuildUpdate models the GUILD_UPDATE object
	GuildUpdate AuditLogEvent = 1
	// ChannelCreate models the CHANNEL_CREATE object
	ChannelCreate AuditLogEvent = 10
	// ChannelUpdate models the CHANNEL_UPDATE object
	ChannelUpdate AuditLogEvent = 11
	// ChannelDelete models the CHANNEL_DELETE object
	ChannelDelete AuditLogEvent = 12
	// ChannelOverwriteCreate models the CHANNEL_OVERWRITE_CREATE object
	ChannelOverwriteCreate AuditLogEvent = 13
	// ChannelOverwriteUpdate models the CHANNEL_OVERWRITE_UPDATE object
	ChannelOverwriteUpdate AuditLogEvent = 14
	// ChannelOverwriteDelete models the CHANNEL_OVERWRITE_DELETE object
	ChannelOverwriteDelete AuditLogEvent = 15
	// MemberKick models the MEMBER_KICK object
	MemberKick AuditLogEvent = 20
	// MemberPrune models the MEMBER_PRUNE object
	MemberPrune AuditLogEvent = 21
	// MemberBanAdd models the MEMBER_BAN_ADD object
	MemberBanAdd AuditLogEvent = 22
	// MemberBanRemove models the MEMBER_BAN_REMOVE object
	MemberBanRemove AuditLogEvent = 23
	// MemberUpdate models the MEMBER_UPDATE object
	MemberUpdate AuditLogEvent = 24
	// MemberRoleUpdate models the MEMBER_ROLE_UPDATE object
	MemberRoleUpdate AuditLogEvent = 25
	// MemberMove models the MEMBER_MOVE object
	MemberMove AuditLogEvent = 26
	// MemberDisconnect models the MEMBER_DISCONNECT object
	MemberDisconnect AuditLogEvent = 27
	// BotAdd models the BOT_ADD object
	BotAdd AuditLogEvent = 28
	// RoleCreate models the ROLE_CREATE object
	RoleCreate AuditLogEvent = 30
	// RoleUpdate models the ROLE_UPDATE object
	RoleUpdate AuditLogEvent = 31
	// RoleDelete models the ROLE_DELETE object
	RoleDelete AuditLogEvent = 32
	// InviteCreate models the INVITE_CREATE object
	InviteCreate AuditLogEvent = 40
	// InviteUpdate models the INVITE_UPDATE object
	InviteUpdate AuditLogEvent = 41
	// InviteDelete models the INVITE_DELETE object
	InviteDelete AuditLogEvent = 42
	// WebhookCreate models the WEBHOOK_CREATE object
	WebhookCreate AuditLogEvent = 50
	// WebhookUpdate models the WEBHOOK_UPDATE object
	WebhookUpdate AuditLogEvent = 51
	// WebhookDelete models the WEBHOOK_DELETE object
	WebhookDelete AuditLogEvent = 52
	// EmojiCreate models the EMOJI_CREATE object
	EmojiCreate AuditLogEvent = 60
	// EmojiUpdate models the EMOJI_UPDATE object
	EmojiUpdate AuditLogEvent = 61
	// EmojiDelete models the EMOJI_DELETE object
	EmojiDelete AuditLogEvent = 62
	// MessageDelete models the MESSAGE_DELETE object
	MessageDelete AuditLogEvent = 72
	// MessageBulkDelete models the MESSAGE_BULK_DELETE object
	MessageBulkDelete AuditLogEvent = 73
	// MessagePin models the MESSAGE_PIN object
	MessagePin AuditLogEvent = 74
	// MessageUnpin models the MESSAGE_UNPIN object
	MessageUnpin AuditLogEvent = 75
	// IntegrationCreate models the INTEGRATION_CREATE object
	IntegrationCreate AuditLogEvent = 80
	// IntegrationUpdate models the INTEGRATION_UPDATE object
	IntegrationUpdate AuditLogEvent = 81
	// IntegrationDelete models the INTEGRATION_DELETE object
	IntegrationDelete AuditLogEvent = 82
	// StageInstanceCreate models the STAGE_INSTANCE_CREATE object
	StageInstanceCreate AuditLogEvent = 83
	// StageInstanceUpdate models the STAGE_INSTANCE_UPDATE object
	StageInstanceUpdate AuditLogEvent = 84
	// StageInstanceDelete models the STAGE_INSTANCE_DELETE object
	StageInstanceDelete AuditLogEvent = 85
)

// AuditLog models the discord audit log object
type AuditLog struct {
	Webhooks        []WebHook       `json:"webhooks"`
	Users           []User          `json:"users"`
	AuditLogEntries []AuditLogEntry `json:"audit_log_entries"`
	Integrations    []Integration   `json:"integrations"`
}

// AuditLogEntry models the discord audit log entry object
type AuditLogEntry struct {
	ID         Snowflake         `json:"id"`
	TargetID   *string           `json:"target_id"`
	Changes    []AuditLogChange  `json:"changes,omitempty"`
	UserID     *Snowflake        `json:"user_id"`
	ActionType AuditLogEvent     `json:"action_type"`
	Options    AuditEntryOptions `json:"options,omitempty"`
	Reason     string            `json:"reason,omitempty"`
}

// AuditEntryOptions models the discord audit entry options
type AuditEntryOptions struct {
	ID               Snowflake `json:"id"`
	DeleteMemberDays string    `json:"delete_member_days"`
	MembersRemoved   string    `json:"members_removed"`
	ChannelID        Snowflake `json:"channel_id"`
	MessageID        Snowflake `json:"message_id"`
	Count            string    `json:"count"`
	Type             string    `json:"type"`
	RoleName         string    `json:"role_name"`
}

// AuditLogChange models the discord audit log change object
type AuditLogChange struct {
	Key      string      `json:"key"`
	NewValue interface{} `json:"new_value,omitempty"`
	OldValue interface{} `json:"old_value,omitempty"`
}
