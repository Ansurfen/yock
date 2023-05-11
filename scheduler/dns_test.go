package scheduler

import (
	"fmt"
	"testing"
)

func TestDNS(t *testing.T) {
	dns := CreateDNS("global.json")
	fmt.Println(dns.GetDriver("yock"))
	fmt.Println(dns.PutDriver("yock", "https://", "https://"))
	fmt.Println(dns.GetDriver("yock"))
	fmt.Println(dns.PutPlugin("yock", "https://", "https://"))
}
