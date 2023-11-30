package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Choose the network device to capture on
	fmt.Println("Enter The Network Device to capture on (by MAC Address): ")
	// Taking input from user
	var selectedDevice string
	fmt.Scanln(&selectedDevice)

	// Open the selected network device for packet capture
	handle, err := pcap.OpenLive(selectedDevice, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("Error opening adapter:", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Capture packets until stopped
	for packet := range packetSource.Packets() {
		// Print the captured packet
		fmt.Println(packet)
	}
}
