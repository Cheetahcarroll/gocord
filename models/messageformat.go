package models

// Format is a string alias that holds the message formats
type Format string

// TimestampStyle is a string alias that holds the Timestamp Styles
type TimestampStyle string

var (
	// UserFormat is the format for a user
	UserFormat = Format("<@%d>")
	// UserNicknameFormat is the format for a user's nickname
	UserNicknameFormat = Format("<@!%d")
	// ChannelFormat is the format for a channel
	ChannelFormat = Format("<#%d>")
	// RoleFormat is the format for a role
	RoleFormat = Format("<@&%d>")
	// CustomEmojiFormat is the format for a custom emoji
	CustomEmojiFormat = Format("<:%s:%v>")
	// CustomAnimatedEmojiFormat is the format for a custom animated emoji
	CustomAnimatedEmojiFormat = Format("<a:%s:%v>")
	// TimestampFormat is the format for a timestamp
	TimestampFormat = Format("<t:%d>")
	// StyledTimestampFormat is the format for a stylized timestamp
	StyledTimestampFormat = Format("t:%d:%s")

	// ShortTime is the style indicator for a short timestamp
	ShortTime = TimestampStyle("t")
	// LongTime is the style indicator for long timestamp
	LongTime = TimestampStyle("T")
	// ShortDate is the style indicator for short date
	ShortDate = TimestampStyle("d")
	// LongDate is the style indicator for long date
	LongDate = TimestampStyle("D")
	// ShortDateTime is the style indicator for short date timestamp
	ShortDateTime = TimestampStyle("f*")
	// LongDateTime is the style indicator for long date timestamp
	LongDateTime = TimestampStyle("F")
	// RelativeTime is the style indicator for relative timestamp
	RelativeTime = TimestampStyle("R")
)
