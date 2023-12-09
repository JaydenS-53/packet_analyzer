# Packet Analyzer
Overview
Packet Analyzer Go is a network packet analysis tool written in Go that captures, filters, and extracts information from network packets. This tool uses the github.com/google/gopacket and github.com/google/gopacket/pcap packages for packet manipulation and capture.

Features
Packet Capture: Capture network packets on a specified network device.
Protocol Support: Analyze TCP and UDP packets.
Custom Filters: Extract specific information from packets.
User Interaction: Allows the user to choose the network device for packet capture.
Dynamic Data Extraction: Extracts relevant data based on the packet type (TCP or UDP).

Prerequisites
Go installed on your machine as well as the following go packages:
github.com/google/gopacket
github.com/google/gopacket/pcap

Install them using:
go get -u github.com/google/gopacket
go get -u github.com/google/gopacket/pcap

Usage

Clone the repository:
git clone https://github.com/your-username/packet_analyzer_go.git

Navigate to the project directory:
cd packet_analyzer_go

Build the application:
go build

Run the application:
./packet_analyzer_go

Enter the network device to capture on when prompted.

Output
The tool captures and filters packets, displaying relevant information based on the packet type (TCP or UDP). The output includes:

Source MAC address, IP address, and port.
DataOffset for TCP packets.
Length for UDP packets.
