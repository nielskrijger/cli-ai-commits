package main

func main() {
	configReader := &FileConfigReader{Filename: "config.yml"}

	apiKey, err := configReader.ReadAPIKey()
	if err != nil {
		panic(err)
	}

	GenerateMessage(apiKey)
}
