package config

import "fmt"

// NewServerBuilder создает новый Builder с параметрами по умолчанию
func NewServerBuilder() *ServerBuilder {
    fmt.Println("[BUILDER] Initializing ServerBuilder with defaults")
    return &ServerBuilder{
        config: &ServerConfig{
            Port:           8090,
            ReadTimeout:    15,
            WriteTimeout:   15,
            EnableLogging:  false,
            MaxHeaderBytes: 1048576, // 1 MB
        },
    }
}
