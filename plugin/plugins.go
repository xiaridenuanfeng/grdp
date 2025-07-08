//go:build !windows
// +build !windows

package plugin

import (
	_ "github.com/tomatome/grdp/plugin/cliprdr"
	_ "github.com/tomatome/grdp/plugin/drdynvc"
	_ "github.com/tomatome/grdp/plugin/rail"
)
