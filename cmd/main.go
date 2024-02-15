package main

import (
	"context"
	"fmt"

	"github.com/pstano1/emailVerifier/config"
)

func main() {
	cfg, err := config.LoadFromPath(context.Background(), "pkl/config.pkl")
	if err != nil {
		panic(err)
	}
	fmt.Printf("I'm running on port: %v\n", cfg.Port)
}
