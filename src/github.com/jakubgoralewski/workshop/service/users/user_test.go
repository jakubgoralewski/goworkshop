package users

import (
	"testing"
)

func TestNewUserTest(t *testing.T) {

	// when
	u := NewUser(1, "jakub", "", 10)

	if u.Id != 1 {
		t.Fail()
	}
}
