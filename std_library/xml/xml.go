package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

type Servers struct {
	XMLName xml.Name `xml:"servers"` //
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

func TestWriteToXml() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{
		ServerName: "Shanghai_VPN",
		ServerIP:   "127.0.0.1",
	})
	v.Svs = append(v.Svs, server{
		ServerName: "Beijing_VPN",
		ServerIP:   "127.0.0.2",
	})

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	//os.Stdout.Write([]byte(xml.Header))
	//os.Stdout.Write(output)

	err = os.WriteFile("servers.xml", output, 0777)
	if err != nil {
		fmt.Println("")
	}
}
