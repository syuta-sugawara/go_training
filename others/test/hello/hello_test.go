package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := GetHello("hogehoge")
	expext := "こんにちは、hogehoge！"
	if result != expext {
		t.Error("\n実際： ", result, "\n理想： ", expext)
	}

	t.Log("TestHello終了")
}
