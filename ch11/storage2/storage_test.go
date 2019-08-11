package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	savedNotifyUser := notifyUser
	savedBytesInUser := bytesInUse

	defer func() {
		notifyUser = savedNotifyUser
		bytesInUse = savedBytesInUser
	}()

	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	// ...simulate a 980MB-used condition...
	bytesInUse = func(username string) int { return quota * 98 / 100 }

	const user = "joe@example.com"
	CheckQuota(user)

	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, wantSubstring)
	}
}
