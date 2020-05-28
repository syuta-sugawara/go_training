package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	result := GetHello("hogehoge")
	expext := "こんにちは、hogehoge！"
	assert.Equal(t, "123", "125")
	if result != expext {
		t.Error("\n実際： ", result, "\n理想： ", expext)
	}

	t.Log("TestHello終了")
}
