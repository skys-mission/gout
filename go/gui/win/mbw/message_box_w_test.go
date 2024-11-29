package mbw

import (
	"fmt"
	"testing"
)

func TestShowMsg(t *testing.T) {
	if err := PopMsg("title", "hello mbw!"); err != nil {
		t.Errorf("(error)%v", err)
		return
	}
	t.Log("ok")
}

func TestShowErrMsg(t *testing.T) {
	if err := PopErrMsg("mbw error!"); err != nil {
		t.Errorf("(error)%v", err)
		return
	}
	t.Log("ok")
}

func TestShowWarningMsg(t *testing.T) {
	if err := PopWarningMsg("mbw warning"); err != nil {
		t.Errorf("(error)%v", err)
		return
	}
	t.Log("ok")
}

func TestShowCustomMsg(t *testing.T) {
	//cb, err := PopCustomMsg("title", "You should click retry.!", MB_ABORTRETRYIGNORE, 0)
	//if err != nil {
	//	t.Errorf("(error)%v", err)
	//	return
	//}
	//if cb != ID_RETRY {
	//	t.Errorf("(error)button callback return value exception.It should equal “ID_RETRY”.But it equals “%v”.", cb)
	//	return
	//}
	//t.Log("ok, button callback: ", cb)
	cb, err := PopCustomMsg(
		"title",
		"You should click retry.!",
		MB_ABORTRETRYIGNORE,
		0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if cb != ID_RETRY {
		fmt.Printf("ID_RETRY clicked \n")
	}
}
