# Ulimit

A pure Go library for managing resource limits.

## Usage

```go
import "github.com/nasdf/ulimit"
```

### Set limits

```go
err := ulimit.SetRlimit(2048)
```

### Get limits

```go
soft, hard, err := ulimit.GetRlimit()
```

## License

MIT