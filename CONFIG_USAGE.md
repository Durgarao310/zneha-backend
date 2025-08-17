# Config Usage Examples

This document shows how to use the configuration system in your application.

## How to Use Config

### 1. Loading Configuration

```go
package main

import (
    "log"
    "github.com/Durgarao310/zneha-backend/internal/config"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // Use configuration
    log.Printf("App: %s v%s", cfg.App.Name, cfg.App.Version)
    log.Printf("Server running on: %s:%s", cfg.Server.Host, cfg.Server.Port)
}
```

### 2. Using Config in Handlers

```go
func NewProductHandler(cfg *config.Config, service service.ProductService) *ProductHandler {
    return &ProductHandler{
        config:  cfg,
        service: service,
    }
}

func (h *ProductHandler) SomeMethod() {
    if h.config.App.Debug {
        log.Println("Debug mode is enabled")
    }
}
```

### 3. Environment-specific Behavior

```go
func setupLogging(cfg *config.Config) {
    if cfg.IsProduction() {
        // Production logging setup
        log.SetLevel(log.WarnLevel)
    } else {
        // Development logging setup
        log.SetLevel(log.DebugLevel)
    }
}
```

## Configuration Priority

1. **Environment Variables** (highest priority)
2. **`.env` file** 
3. **Default values** (lowest priority)

## Environment Variables

All config values can be overridden with environment variables:

```bash
export PORT=3000
export DB_HOST=production-db.example.com
export JWT_SECRET=super-secret-production-key
export ENVIRONMENT=production
```
