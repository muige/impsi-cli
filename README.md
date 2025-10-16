# impsi-cli

A terminal-based unit converter for SI ↔ Imperial conversions.

## Installation

```bash
go build
```

## Usage

```bash
./impsi-cli
```

Enter a number to see real-time conversions including:
- Distance: miles, feet, inches, yards ↔ km, meters, cm
- Weight: pounds, ounces ↔ kg, grams
- Volume: gallons, fl oz ↔ liters, mL

Press Ctrl-C or Esc to quit.

## Development

- Build: `go build`
- Run: `go run main.go`
- Test: `go test ./...`
