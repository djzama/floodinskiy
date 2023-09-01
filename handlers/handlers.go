package handlers

import (
	"floodinskiy/packets"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
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

	packet := packets.NewPacket(protocol, targetIP, targetPort, packetSize)

	for i := 0; i < packetCount; i++ {
		err := packet.Send()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error sending packet: %s", err), http.StatusInternalServerError)
		}
		logMessage := fmt.Sprintf("%s:%d:%s:%d", protocol, i+1, targetIP, targetPort)
		fmt.Fprint(w, logMessage)
		time.Sleep(100 * time.Millisecond)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
