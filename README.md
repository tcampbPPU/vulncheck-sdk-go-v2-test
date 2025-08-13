# VulnCheck Go SDK V2 Example Playground

## CLI Usage

This project includes a lightweight CLI interface to easily test different SDK functions without modifying code.

### Commands

```bash
# List all available test functions
go run main.go list

# Run a specific function
go run main.go run <function-name>
go run main.go <function-name>  # shorthand

# Show help
go run main.go help
```

### Available Functions

Use `go run main.go list` to see all available functions, including:

- `index-initial-access` - Get Index Initial Access
- `index-vulnrichment` - Get Index Vulnrichment
- `index-cve-filter` - Get Index with CVE Filter
- `index-botnet-filter` - Get Index with Botnet Filter
- `index-ip-intel` - Get Index IP Intel
- `browse-indexes` - Browse Indexes
- `browse-backups` - Browse Backups
- `backup` - Get Index Backup
- `purl` - Get PURL
- `rule` - Get Rule
- `tag` - Get Tag
- `pdns` - Get PDNS
- `cpe` - Get CPE

### Makefile Shortcuts

```bash
# List functions
make list

# Run specific function
make run FUNC=browse-indexes

# Show help
make help
```

## Testing

The project includes comprehensive tests for the CLI functionality with 100% code coverage.

### Running Tests

```bash
# Run all tests
make test

# Run tests with verbose output
make test-verbose

# Run tests with coverage information
make test-coverage

# Run CLI-specific tests
make test-cli

# Run benchmark tests
make test-bench
```

### Test Coverage

- **CLI Package**: 100% statement coverage
- **Integration Tests**: End-to-end testing with actual binary execution
- **Performance Tests**: Benchmarks for critical functions
- **Mock Testing**: Isolated testing without external API calls

See [TESTING.md](TESTING.md) for detailed testing documentation.

### Examples

```bash
# List all available functions
go run main.go list

# Run the browse-indexes function
go run main.go browse-indexes

# Run using the run command
go run main.go run tag

# Using Makefile
make run FUNC=browse-indexes
```