// This example uses gstreamer's device provider api.
//
// https://gstreamer.freedesktop.org/documentation/gstreamer/gstdeviceprovider.html
package main

import (
	"fmt"
	"os"

	"github.com/go-gst/go-glib/glib"
	"github.com/rdkbob/go-gst/examples"
	"github.com/rdkbob/go-gst/gst"
)

func runPipeline(loop *glib.MainLoop) error {

	gst.Init(nil)
	fmt.Println("Running device provider")
	// if len(os.Args) < 2 {
	// 	fmt.Printf("USAGE: %s <uri>\n", os.Args[0])
	// 	os.Exit(1)
	// }

	// uri := os.Args[1]
	fmt.Println("Creating device monitor")

	// provider := gst.FindDeviceProviderByName("foo")
	// fmt.Println("Created device provider", provider)

	provider := gst.FindDeviceProviderByName("avfdeviceprovider")
	fmt.Println("Created device provider", provider)

	if provider == nil {
		fmt.Println("No provider found")
		os.Exit(2)
	}

	fmt.Println("Getting device provider bus")
	bus := provider.GetBus()
	fmt.Println("Got device provider bus", bus)

	bus.AddWatch(func(msg *gst.Message) bool {
		switch msg.Type() {
		case gst.MessageDeviceAdded:
			message := msg.ParseDeviceAdded().GetDisplayName()
			fmt.Println("Added: ", message)
		case gst.MessageDeviceRemoved:
			message := msg.ParseDeviceRemoved().GetDisplayName()
			fmt.Println("Removed: ", message)
		default:
			// All messages implement a Stringer. However, this is
			// typically an expensive thing to do and should be avoided.
			fmt.Println("Type: ", msg.Type())
			fmt.Println("Message: ", msg)
		}
		return true
	})

	fmt.Println("Starting device monitor")
	provider.Start()
	fmt.Println("Started device monitor")

	fmt.Println("listing devices from provider")
	devices := provider.GetDevices()
	for i, v := range devices {
		fmt.Printf("Device: %d %s\n", i, v.GetDisplayName())
	}

	loop.Run()

	return nil
}

func main() {
	examples.RunLoop(func(loop *glib.MainLoop) error {
		return runPipeline(loop)
	})
}
