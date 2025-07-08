//go:build !windows
// +build !windows

package cliprdr

const (
	CLIPRDR_SVC_CHANNEL_NAME     = "cliprdr"
	CHANNEL_OPTION_INITIALIZED   = 0
	CHANNEL_OPTION_ENCRYPT_RDP   = 1
	CHANNEL_OPTION_COMPRESS_RDP  = 2
	CHANNEL_OPTION_SHOW_PROTOCOL = 3
	FILE_ATTRIBUTE_DIRECTORY     = 0x10
)

type CliprdrClient struct{}

func (c *CliprdrClient) withOpenClipboard(f func()) {}

func EmptyClipboard() {}

type Control struct{}
type ClipWatcher struct{}