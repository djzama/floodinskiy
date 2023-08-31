package packets

import (
	"fmt"
	"net"
)

func sendTCPPacket(targetIP string, targetPort int, packetSize int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", targetIP, targetPort))
	if err != nil {
		fmt.Println("TCP connection error:", err)
		return
	}
	defer conn.Close()

	data := make([]byte, packetSize)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("TCP write error:", err)
	}
}

func sendUDPPacket(targetIP string, targetPort int, packetSize int) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", targetIP, targetPort))
	if err != nil {
		fmt.Println("UDP connection error:", err)
		return
	}
	defer conn.Close()

	data := make([]byte, packetSize)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("UDP write error:", err)
	}
}

func sendICMPPacket(targetIP string, packetSize int) {
	conn, err := net.Dial("ip4:icmp", targetIP)
	if err != nil {
		fmt.Println("ICMP connection error:", err)
		return
	}
	defer conn.Close()

	data := make([]byte, packetSize)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("ICMP write error:", err)
	}
}

func sendPOPPacket(targetIP string, targetPort int, packetSize int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", targetIP, targetPort))
	if err != nil {
		fmt.Println("POP3 connection error:", err)
		return
	}
	defer conn.Close()

	data := make([]byte, packetSize)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("POP3 write error:", err)
	}
}

func sendSMTPPacket(targetIP string, targetPort int, packetSize int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", targetIP, targetPort))
	if err != nil {
		fmt.Println("SMTP connection error:", err)
		return
	}
	defer conn.Close()

	data := make([]byte, packetSize)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("SMTP write error:", err)
	}
}

func sendDNSPacket(targetIP string, targetPort int, packetSize int) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", targetIP, targetPort))
	if err != nil {
		fmt.Println("DNS connection error:", err)
		return
	}
	defer conn.Close()

	data := make([]byte, packetSize)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("DNS write error:", err)
	}
}
