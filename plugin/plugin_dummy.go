
//go:build !windows
// +build !windows

package plugin

type Plugin interface {
    Name() string
    Init() error
    Handle(data []byte) ([]byte, error)
}

// 空实现的结构体（可选）
type DummyPlugin struct{}

func (d *DummyPlugin) Name() string                  { return "dummy" }
func (d *DummyPlugin) Init() error                   { return nil }
func (d *DummyPlugin) Handle(data []byte) ([]byte, error) { return nil, nil }
