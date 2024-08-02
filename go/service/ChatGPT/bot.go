package ChatGPT

import "fmt"

func bot(question string) string {
	config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	transport, err := SetupProxy(config.ProxyOption)
	if err != nil {
		fmt.Println("Error setting up proxy:", err)
		return
	}

	_, err := CallAPI(question, config, transport)
	if err != nil {
		fmt.Println("Error calling API:", err)
		continue
	}

	fmt.Println()
}
