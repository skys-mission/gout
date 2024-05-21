package mbw

import (
	"testing"
)

func TestWinMsg(t *testing.T) {
	if err := WinMsg("title", "hello mbw!"); err != nil {
		t.Errorf("(error)%v", err)
		return
	}
	t.Log("ok")
}

func TestWinErrMsg(t *testing.T) {
	if err := WinErrMsg("mbw error!"); err != nil {
		t.Errorf("(error)%v", err)
		return
	}
	t.Log("ok")
}

func TestWinCustomMsg(t *testing.T) {
	cb, err := WinCustomMsg("title", "You should click retry.!", MB_ABORTRETRYIGNORE, 0)
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
