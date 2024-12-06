package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Main() {
	handlerDir := "./internal/core/handlers"
	usecaseDir := "./internal/core/usecases"
	repositoryDir := "./internal/core/repositories"

	_, err := exec.LookPath("mockery")
	if err != nil {
		fmt.Println("Mockery could not be found, please install it first")
		os.Exit(1)
	}

	generateMocks := func(dir string) error {
		fmt.Printf("Generating mocks for %s\n", dir)

		cmd := exec.Command("mockery", "--all", "--dir", dir)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error generating mocks for %s: %s\n", dir, err)
			return err
		}

		return nil
	}

	for _, dir := range []string{handlerDir, usecaseDir, repositoryDir} {
		if err := generateMocks(dir); err != nil {
			fmt.Printf("An error occurred while generating mocks: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Mocks generation completed.")
}
