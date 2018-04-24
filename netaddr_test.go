package netaddr

import (
	"testing"
)

func TestIPnetworkCounts(t *testing.T) {
	// block := "192.0.2.16/29"
	tests := []struct {
		Name     string
		CIDR     string
		Expected int
	}{
		{
			Name:     "Twitter Example",
			CIDR:     "192.168.0.16/29",
			Expected: 8,
		}, {
			Name:     "one address",
			CIDR:     "192.168.0.1/32",
			Expected: 1,
		}, {
			Name:     "slash 31?",
			CIDR:     "192.168.0.1/31",
			Expected: 2,
		}, {
			Name:     "classic /24",
			CIDR:     "192.168.0.1/24",
			Expected: 256,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			addrs, err := IPNetwork(test.CIDR)

			if err != nil {
				t.Fatalf("IPNetwork(%q) failed with %v", test.CIDR, err)
			}

			t.Logf("%s: %v", test.Name, addrs)
			if len(addrs) != test.Expected {
				t.Fatalf("IPNetwork(%s) returned %d addresses; expected %d", test.CIDR, len(addrs), test.Expected)

			}
		})
	}
}
