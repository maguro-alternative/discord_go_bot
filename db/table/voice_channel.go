package table

type VcChannel struct {
	VcID            string   `db:"vc_id"`
	GuildID         string   `db:"guild_id"`
	SendSignal      bool     `db:"send_signal"`
	SendChannelID   string   `db:"send_channel_id"`
	JoinBot         bool     `db:"join_bot"`
	EveryoneMention bool     `db:"everyone_mention"`
	MentionRoleIDs  []string `db:"mention_role_ids"`
}
