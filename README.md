# Logger

Универсальный пакет для логирования с поддержкой различных бэкендов и gRPC.

## Особенности

- Единый интерфейс для разных логгеров
- Поддержка zap
- Встроенная поддержка gRPC логирования
- Структурированное логирование

## Установка

```bash
go get github.com/peso4eeek/logger
```

## Использование

### Базовое использование

```go
import (
    "github.com/peso4eeek/logger"
    "github.com/peso4eeek/logger/zapadapter"
    "go.uber.org/zap"
)

func main() {
    zapLogger, _ := zap.NewProduction()
    log := zapadapter.NewZapAdapter(zapLogger)
    
    log.Info("Приложение запущено", "version", "1.0.0")
}
```

### Использование с gRPC

```go
import (
    "github.com/peso4eeek/logger"
    "google.golang.org/grpc"
)

func main() {
    // ... инициализация логгера ...
    
    cc, err := grpc.Dial(
        "localhost:50051",
        grpc.WithChainUnaryInterceptor(
            grpclog.UnaryClientInterceptor(logger.GRPCInterceptor(log)),
        ),
    )
}
```

## Адаптеры

### Zap

```go
import (
    "github.com/peso4eeek/logger/zapadapter"
    "go.uber.org/zap"
)

zapLogger, _ := zap.NewProduction()
log := zapadapter.NewZapAdapter(zapLogger)
```

## Лицензия

MIT 