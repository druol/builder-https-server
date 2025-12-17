package config

import "fmt"

// ServerConfig хранит параметры HTTP сервера
type ServerConfig struct {
    Port            int
    ReadTimeout     int
    WriteTimeout    int
    EnableLogging   bool
    MaxHeaderBytes  int
}

// Print выводит конфигурацию в консоль
func (sc *ServerConfig) Print() {
    fmt.Println("[CONFIG] Server Configuration:")
    fmt.Printf("[CONFIG]   Port: %d\n", sc.Port)
    fmt.Printf("[CONFIG]   ReadTimeout: %ds\n", sc.ReadTimeout)
    fmt.Printf("[CONFIG]   WriteTimeout: %ds\n", sc.WriteTimeout)
    
    loggingStatus := "disabled"
    if sc.EnableLogging {
        loggingStatus = "enabled"
    }
    fmt.Printf("[CONFIG]   Logging: %s\n", loggingStatus)
    fmt.Printf("[CONFIG]   MaxHeaderBytes: %d\n", sc.MaxHeaderBytes)
}
