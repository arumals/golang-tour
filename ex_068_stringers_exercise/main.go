package main

import (
	"fmt"
)

type IpAddress [4]int

func (ipa IpAddress) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipa[0], ipa[1], ipa[2], ipa[3])
}

func main() {
	var ip IpAddress = IpAddress{8, 8, 8, 8}
	fmt.Println(ip)
}
