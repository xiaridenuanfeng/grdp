//go:build !windows
// +build !windows

package rail

const (
	RAIL_SVC_CHANNEL_NAME        = "rail"
	CHANNEL_OPTION_INITIALIZED   = 0
	CHANNEL_OPTION_ENCRYPT_RDP   = 1
	CHANNEL_OPTION_COMPRESS_RDP  = 2
	CHANNEL_OPTION_SHOW_PROTOCOL = 3
)