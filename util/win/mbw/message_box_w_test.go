package mbw

import (
	"testing"
)

func TestShowMsg(t *testing.T) {
	if err := ShowMsg("title", "hello mbw!"); err != nil {
		t.Errorf("(error)%v", err)
		return
	}
	t.Log("ok")
}

func TestShowErrMsg(t *testing.T) {
	if err := ShowErrMsg("mbw error!"); err != nil {
		t.Errorf("(error)%v", err)
		return
	}
	t.Log("ok")
}

func TestShowCustomMsg(t *testing.T) {
	cb, err := ShowCustomMsg("title", "You should click retry.!", MB_ABORTRETRYIGNORE, 0)
	if err != nil {
		t.Errorf("(error)%v", err)
		return
	}
	if cb != ID_RETRY {
		t.Errorf("(error)button callback return value exception.It should equal “ID_RETRY”.But it equals “%v”.", cb)
		return
	}
	t.Log("ok, button callback: ", cb)
}
