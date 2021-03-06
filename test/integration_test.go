package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendStatusLastActivity(t *testing.T) {
	resp := integration.SendLastActivity("123456", "فروغ")
	m := fmt.Sprintf("تسک %s از وضعیت %s به وضعیت %s تغییر کرد.", "Done", "Doing", "github")
	expected := fmt.Sprintf(`{"text":"%s"}`, m)
	assert.Equal(t, resp, expected)
}
