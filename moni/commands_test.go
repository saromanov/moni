package moni

import (
	"testing"
)

func TestDiskspace(t *testing.T) {
	_, err := diskSpace(" ")
	if err == nil {
		t.Error("Expected to getting error")
	}

	_, erre := diskSpace("123456")
	if erre == nil {
		t.Error("Expected to getting error")
	}

	value, err2 := diskSpace(`Filesystem              Size  Used Avail Use% Mounted on
/dev/sda1                87G   73G  9.7G  89% /
none                    4.0K     0  4.0K   0% /none
none                    488M   12K  488M   1% /none
none                   101M  1.4M   99M   2% /none
none                    5.0M     0  5.0M   0% /none
none                    501M  788K  501M   1% /none
none                    100M  144K  100M   1% /none
123   87G   73G  9.7G  89% /none/none
`)
	if err2 != nil {
		t.Error("Error in disk space")
	}

	if value != "9.7G" {
		t.Errorf("Expected %s, found %s", value, "9.7G")
	}

}
