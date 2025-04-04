package main

import (
	"fmt"
)

type DNSHeader struct {
	TransactionID  uint16
	Flags          uint16
	NumQuestions   uint16
	NumAnswers     uint16
	NumAuthorities uint16
	NumAdditionals uint16
}

type DNSResourceRecord struct {
	DomainName         string
	Type               uint16
	Class              uint16
	TimeToLive         uint32
	ResourceDataLength uint16
	ResourceData       []byte
}

// Type and Class values for DNSResourceRecord
const (
	TypeA                  uint16 = 1 // a host address
	ClassINET              uint16 = 1 // the Internet
	FlagResponse           uint16 = 1 << 15
	UDPMaxMessageSizeBytes uint   = 512 // RFC1035
)

func main() {
	fmt.Print("Test 1")
}
