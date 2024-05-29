package privilegew

import "testing"

func TestEnableDebugPrivilege(t *testing.T) {
	if err := EnableDebugPrivilege(); err != nil {
		t.Errorf("err: %v", err)
		return
	}
	t.Log("ok")
}
