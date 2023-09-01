package packets

import (
	"fmt"
	"net"
)

type Packet struct {
	Protocol   string
	TargetIP   string
	TargetPort int
	PacketSize int
}

func NewPacket(protocol, targetIP string, targetPort, packetSize int) *Packet {
	return &Packet{
		Protocol:   protocol,
		TargetIP:   targetIP,
		TargetPort: targetPort,
		PacketSize: packetSize,
	}
}

func (p *Packet) Send() error {
	switch p.Protocol {
	case "tcp":
		return p.sendTCPPacket()
	case "udp":
		return p.sendUDPPacket()
	case "icmp":
		return p.sendICMPPacket()
	case "pop3":
		return p.sendPOPPacket()
	case "smtp":
		return p.sendSMTPPacket()
	case "dns":
		return p.sendDNSPacket()
	default:
		return fmt.Errorf("Unsupported protocol: %s", p.Protocol)
	}
}

func (p *Packet) sendTCPPacket() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", p.TargetIP, p.TargetPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	data := make([]byte, p.PacketSize)
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (p *Packet) sendUDPPacket() error {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", p.TargetIP, p.TargetPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	data := make([]byte, p.PacketSize)
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (p *Packet) sendICMPPacket() error {
	conn, err := net.Dial("ip4:icmp", p.TargetIP)
	if err != nil {
		return err
	}
	defer conn.Close()

	data := make([]byte, p.PacketSize)
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (p *Packet) sendPOPPacket() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", p.TargetIP, p.TargetPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	data := make([]byte, p.PacketSize)
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (p *Packet) sendSMTPPacket() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", p.TargetIP, p.TargetPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	data := make([]byte, p.PacketSize)
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (p *Packet) sendDNSPacket() error {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", p.TargetIP, p.TargetPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	data := make([]byte, p.PacketSize)
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}
