// handlers/handlers.go

package handlers

import (
	packets "floodinskiy/packet"
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	protocol := r.FormValue("protocol")
	targetIP := r.FormValue("targetIP")
	targetPort, _ := strconv.Atoi(r.FormValue("targetPort"))
	packetSize, _ := strconv.Atoi(r.FormValue("packetSize"))
	packetCount, _ := strconv.Atoi(r.FormValue("packetCount"))

	for i := 0; i < packetCount; i++ {
		switch protocol {
		case "tcp":
			packets.sendTCPPacket(targetIP, targetPort, packetSize)
		case "udp":
			packets.sendUDPPacket(targetIP, targetPort, packetSize)
		case "icmp":
			packets.SendICMPPacket(targetIP, packetSize)
		case "pop3":
			packets.SendPOPPacket(targetIP, targetPort, packetSize)
		case "smtp":
			packets.SendSMTPPacket(targetIP, targetPort, packetSize)
		case "dns":
			packets.SendDNSPacket(targetIP, targetPort, packetSize)
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
