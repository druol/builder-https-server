package config

import "fmt"

// ServerBuilder реализует паттерн Builder
type ServerBuilder struct {
    config *ServerConfig
}

// SetPort устанавливает порт сервера
func (sb *ServerBuilder) SetPort(port int) *ServerBuilder {
    sb.config.Port = port
    fmt.Printf("[BUILDER] Port set to %d\n", port)
    return sb
}

// SetReadTimeout устанавливает таймаут чтения
func (sb *ServerBuilder) SetReadTimeout(seconds int) *ServerBuilder {
    sb.config.ReadTimeout = seconds
    fmt.Printf("[BUILDER] ReadTimeout set to %ds\n", seconds)
    return sb
}

// SetWriteTimeout устанавливает таймаут записи
func (sb *ServerBuilder) SetWriteTimeout(seconds int) *ServerBuilder {
    sb.config.WriteTimeout = seconds
    fmt.Printf("[BUILDER] WriteTimeout set to %ds\n", seconds)
    return sb
}

// EnableLogging включает логирование
func (sb *ServerBuilder) EnableLogging() *ServerBuilder {
    sb.config.EnableLogging = true
    fmt.Println("[BUILDER] Logging enabled")
    return sb
}

// SetMaxHeaderBytes устанавливает максимальный размер заголовков
func (sb *ServerBuilder) SetMaxHeaderBytes(bytes int) *ServerBuilder {
    sb.config.MaxHeaderBytes = bytes
    fmt.Printf("[BUILDER] MaxHeaderBytes set to %d\n", bytes)
    return sb
}

// Build завершает построение и возвращает конфигурацию
func (sb *ServerBuilder) Build() *ServerConfig {
    fmt.Println("[BUILDER] Configuration built successfully")
    return sb.config
}
