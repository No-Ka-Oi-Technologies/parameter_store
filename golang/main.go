package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	// Get the path from command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [path]")
		return
	}
	path := os.Args[1]

	// Load AWS configuration from the environment or shared configuration file
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Println("Error loading AWS configuration: ", err)
		return
	}

	// Create an SSM client with the configuration
	svc := ssm.NewFromConfig(cfg)

	// Set up parameters for the request
	input := &ssm.GetParametersByPathInput{
		Path:           aws.String(path),
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	}

	// Create a map to store parameter names and values
	params := make(map[string]string)

	// Retrieve parameters from the parameter store
	for {
		result, err := svc.GetParametersByPath(context.Background(), input)
		if err != nil {
			fmt.Println("Error retrieving parameters: ", err)
			return
		}

		// Add each parameter to the map
		for _, p := range result.Parameters {
			paramName := filepath.Base(*p.Name)
			params[paramName] = *p.Value
		}

		// If there are no more parameters, break out of the loop
		if result.NextToken == nil {
			break
		}

		// Set the next token to retrieve the next batch of parameters
		input.NextToken = result.NextToken
	}

	// Sort the parameter names alphabetically
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create or open the .env file for writing
	file, err := os.Create(".env")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	// Write each parameter name and value to the file in alphabetical order
	for _, k := range keys {
		line := fmt.Sprintf("%s=%s\n", k, params[k])
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file: ", err)
			return
		}
	}

	fmt.Println("Parameters written to .env")
}
