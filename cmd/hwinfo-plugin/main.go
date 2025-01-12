package main

import (
	"log"

	"github.com/hashicorp/go-plugin"
	hwinfoplugin "github.com/shayne/go-hwinfo-streamdeck-plugin/internal/hwinfo/plugin"
	hwsensorsservice "github.com/shayne/go-hwinfo-streamdeck-plugin/pkg/service"
)

func main() {
	service := hwinfoplugin.StartService()
	go func() {
		for {
			err := service.Recv()
			if err != nil {
				log.Printf("service recv failed: %v\n", err)
			}
		}
	}()

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hwsensorsservice.Handshake,
		Plugins: map[string]plugin.Plugin{
			"hwinfoplugin": &hwsensorsservice.HardwareServicePlugin{Impl: &hwinfoplugin.Plugin{Service: service}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
