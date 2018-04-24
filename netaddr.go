package netaddr

import (
	"net"
)

// IPNetwork returns a slice of net.IP addresses that fall within a range
// spcified by a cidr notation block.
func IPNetwork(block string) ([]net.IP, error) {
	ip, netw, err := net.ParseCIDR(block)

	if err != nil {
		return nil, err
	}

	// The net.IP and net.IPMask types store addresses and masks as bytes.
	// btoi() pakcs a series of bytes into one int32 so that we can use bitwise
	// operations to determine the network address and range for a CIDR network
	// range.
	//
	// itob() takes an int32 and separates it into bytes.
	btoi := func(b1, b2, b3, b4 byte) int32 {
		return int32(b1)<<24 | int32(b2)<<16 | int32(b3)<<8 | int32(b4)
	}
	itob := func(i int32) []byte {
		return []byte{byte(i >> 24 & 0xff), byte(i >> 16 & 0xff), byte(i >> 8 & 0xff), byte(i & 0xff)}
	}

	addr := btoi(ip[12], ip[13], ip[14], ip[15])
	mask := ^btoi(netw.Mask[0], netw.Mask[1], netw.Mask[2], netw.Mask[3])

	bcast := addr | mask
	network := bcast - mask

	var rc []net.IP
	// we treat the network addresses as int32 and simply iterate from bcast to
	// network to find addresses within the CIDR block.
	for i := network; i <= bcast; i++ {
		// convert the int32 address to a net.IP and append it to our return value.
		rc = append(rc, net.IP(itob(i)))
	}

	return rc, nil
}
