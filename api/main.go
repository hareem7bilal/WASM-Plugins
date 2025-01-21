package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/extism/go-sdk"
	"os"
)

func main() {

	// Create a manifest for the plugin
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: "./plugin/plugin.wasm", // Path to WASM file
			},
		},
		AllowedHosts: []string{"*"},
	}

	// Initialize the plugin
	ctx := context.Background()
	config := extism.PluginConfig{EnableWasi: true}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})
	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	defer plugin.Close()

	// Input for the plugin function

	apiRequest := map[string]interface{}{
		"url":    "https://jsonplaceholder.typicode.com/posts",
		"method": "POST",
		"headers": map[string]string{
			"Content-Type": "application/json",
		},
		"body": `{"title": "foo", "body": "bar", "userId": 1}`,
	}

	// apiRequest := map[string]interface{}{
	// 	"url":    "https://jsonplaceholder.typicode.com/posts",
	// 	"method": "GET",
	// 	"headers": map[string]string{
	// 		"Accept": "application/json",
	// 	},
	// 	"body": "",
	// }

	// apiRequest := map[string]interface{}{
	// 	"url":    "http://localhost:3100/mdm-api/graphql",
	// 	"method": "POST",
	// 	"headers": map[string]string{
	// 		"Content-Type": "application/json",
	// 	},
	// 	"body": `{"query": "query { getAllDataObjects(pagination: {offset: 0, limit: 5}, QueryInput:{filterInputString:\"{}\"}) { info } }"}`,
	// }

	// apiRequest := map[string]interface{}{
	// 	"url":    "http://localhost:3100/mdm-api/graphql",
	// 	"method": "POST",
	// 	"headers": map[string]string{
	// 		"Content-Type": "application/json",
	// 	},
	// 	"body": `{"query": "query { getAllDataObjects(pagination: {offset: 0, limit: 5}, QueryInput:{filterInputString:\"{}\"}) { info } }"}`,
	// }

	// apiRequest := map[string]interface{}{
	// 	"url":     "http://localhost:3100/mdm-api/graphql",
	// 	"method":  "POST",
	// 	"headers": map[string]string{"Content-Type": "application/json", "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjIjoiZmU5Y2NiNjctNjRiYi00YjQ3LTgwY2YtYTBjYWEzMGFiOWQ0IiwidSI6ImhhcmVlbS5iaWxhbEB6b25lcy5jb20iLCJzdWIiOjEsIm0iOnsicm9sZXMiOlszLDMsMywzXSwibmFtZSI6IkhhcmVlbSBCaWxhbCIsIm9pZCI6bnVsbCwiYXBwcyI6WzEsMiwzLDRdfSwiaWF0IjoxNzMwOTc4MzEyLCJleHAiOjE3MzEwNjQ3MTJ9.H2eyitHXJHJWsYb8Ejtj6lywcybvansRo1MR_SBh7JA"},
	// 	"body":    `{"query": "mutation { manageDataObject(dataObjectInput: { id: \"10033af0-e61a-4b6a-8d2b-00fa18cd2a05\", name: \"CSDataObject6350\", tableMetadata: [ { id: \"fee6d800-31d2-4c2c-ba39-12702cbbcd0e\", name: \"flow1tb11wwrgddwfwfwew15\", tableAttributes: [ { id: \"48c9b65b-5474-48e0-b369-6d532de4a951\", name: \"id\", dataType: \"uuid\" } ] } ] }) { status message dataObject { id tableMetadata { id tableAttributes { id name } } } } }"}`,
	// }

	// apiRequest := map[string]interface{}{
	// 	"appUrl": "http://localhost:3100/mdm-api/graphql",
	// 	"method":     "POST",
	// 	"headers": map[string]string{
	// 		"Content-Type": "application/json",
	// 	},
	// 	"body": `{"query": "query { getAllDataObjects(pagination: {offset: 0, limit: 5}, QueryInput:{filterInputString:\"{}\"}) { info } }"}`,
	// }

	// Convert request to JSON
	input, err := json.Marshal(apiRequest)
	if err != nil {
		fmt.Printf("Failed to marshal request: %v\n", err)
		os.Exit(1)
	}

	// Call the exported function (makeRequest) in your WASM plugin
	exitCode, output, err := plugin.Call("makeRequest", input)
	if err != nil {
		fmt.Printf("Plugin execution failed: %v (exit code: %d)\n", err, exitCode)
		os.Exit(1)
	}

	// Print the output
	fmt.Printf("Plugin output: %s\n", string(output))
}
