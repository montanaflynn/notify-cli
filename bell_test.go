package notify

import "testing"

func TestBellNotifier(t *testing.T) {
	bell := bellNotifier{}
	err := bell.Notify()
	if err != nil {
		t.Error(err)
	}
}
