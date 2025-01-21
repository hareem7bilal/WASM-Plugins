package main

import (
    "context"
    "fmt"
    "github.com/extism/go-sdk"
    "os"
)

func main() {
    // Create a manifest for the plugin
    manifest := extism.Manifest{
        Wasm: []extism.Wasm{
            extism.WasmFile{
                Path: "./plugin/plugin.wasm",
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

    // Call the exported function `httpGet`
    exitCode, output, err := plugin.Call("httpGet", []byte{})
    if err != nil {
        fmt.Printf("Plugin execution failed: %v (exit code: %d)\n", err, exitCode)
        os.Exit(1)
    }

    // Print the output
    fmt.Printf("Plugin output: %s\n", string(output))
}