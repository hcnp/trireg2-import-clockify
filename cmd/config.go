package cmd

import "os"

type config struct {
  apiKey string
}

func loadEnvVars() {
  os.Getenv("CLOCKIFY_API_KEY")
}
