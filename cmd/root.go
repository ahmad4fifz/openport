package cmd

import (
	"log"
	"net"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "openport",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")

		// Check valid port number supplied
		if PortValid(port) {
			// Listening on port
			listener, err := net.Listen("tcp4", ":"+PortToString(port))
			if err != nil {
				log.Printf("Error: Can't listen on port %d: %s", port, err)
				os.Exit(1)
			}
			// Close the listener when the application closes.
			defer listener.Close()
			log.Printf("Listening on port %d", port)
			for {
				// Listen for an incoming connection.
				conn, err := listener.Accept()
				if err != nil {
					log.Printf("Error accepting: %s", err)
					os.Exit(1)
				}
				// Display the remote address
				log.Printf("Received connection from %s", conn.RemoteAddr())
			}
		} else {
			log.Printf("Error: Not a valid port: %d", port)
		}
	},
}

// Convert port number form data type int to string
func PortToString(port int) (portString string) {
	portString = strconv.Itoa(port)
	return
}

// Validate port number
func PortValid(port int) bool {
	portString := PortToString(port)
	// Regex matching port number
	r, _ := regexp.Compile("^()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])$")
	return r.Match([]byte(portString))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Define flag(s)
	rootCmd.PersistentFlags().IntP("port", "p", 0, "tcp port: 1-65535")

	// Set flag(s) as Required
	rootCmd.MarkPersistentFlagRequired("port")
}
