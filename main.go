package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// function to get data between two strings, inlcuding start string but not end string
func getDataBetweenStrings(fullString, startString, endString string) string {
	// Find the starting index of the startString
	startIndex := strings.Index(fullString, startString)
	if startIndex == -1 {
		return ""
	}

	// Find the ending index of the endString
	endIndex := strings.Index(fullString, endString)
	if endIndex == -1 {
		return ""
	}

	// Extract the substring between start and end indices, including start string but excluding end string
	result := fullString[startIndex:endIndex]

	return result
}

// function to filter data from packet
func filter_packet(packet string) {
	var dstPort string
	if strings.Contains(packet, "TCP") {
		dstPort = "DataOffset"
		println("TCP Packet: ")

	} else if strings.Contains(packet, "UDP") {
		dstPort = "Length"
		println("UDP Packet: ")
	}
	fmt.Println(getDataBetweenStrings(packet, "SrcMAC", "EthernetType"))
	fmt.Println(getDataBetweenStrings(packet, "SrcIP", "Options"))
	fmt.Println(getDataBetweenStrings(packet, "SrcPort", dstPort))
}

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
		// Convert packet data type to string
		strPacket := packet.String()
		// filter packet and output
		filter_packet(strPacket)
		println("\n")
	}
}
