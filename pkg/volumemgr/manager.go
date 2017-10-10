package volumemgr

import "os"

type Manager struct {
	Dir          string
	MountOptions map[string]string
}

func (m *Manager) Start() error {
	os.MkdirAll(m.Dir, 0744)
	return nil
}

func (m *Manager) Stop() error {
	os.RemoveAll(m.Dir)
	return nil
}
