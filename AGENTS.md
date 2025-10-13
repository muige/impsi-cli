# Agent Guidelines for impsi

## Build/Test Commands
- Build: `go build`
- Run: `go run main.go`
- Test all: `go test ./...`
- Test package: `go test ./conversions`
- Format: `go fmt ./...`
- Lint: `go vet ./...`

## Code Style
- **Imports**: Group stdlib first, then third-party, then local packages (see main.go:3-13)
- **Naming**: Use camelCase for unexported, PascalCase for exported; short variable names (m, s, ti)
- **Types**: Define structs with clear field comments (see conversionPair in main.go:29-40)
- **Functions**: Return functions as closures when needed (e.g., FtToM() returns func(float64) float64)
- **Comments**: Package comments required; function comments for complex logic; inline for clarity
- **Error handling**: Return errors, check with if err != nil, exit with os.Exit(1) in main
- **Line length**: Keep reasonable (~100 chars)
- **Variables**: Use descriptive names for globals (conversions, gap), short for locals (s, m, r, i)
