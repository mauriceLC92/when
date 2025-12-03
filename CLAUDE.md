# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run all tests
go test ./...

# Run tests for a specific language package
go test ./rules/en/...
go test ./rules/ru/...

# Run a specific test
go test ./rules/en -run TestDeadline

# Run tests with verbose output
go test -v ./...
```

## Architecture

This is a Go library for parsing natural language date/time expressions. The parser uses a rule-based system with pluggable rules per language.

### Core Components

- `when.go` - Main `Parser` struct with `Parse()` method that coordinates rule matching and application
- `rules/rules.go` - `Rule` interface and `Match` struct. Rules use regex patterns via the `F` struct which implements `Find()`
- `rules/context.go` - `Context` accumulates parsed time components (Year, Month, Day, Hour, etc.) and relative durations

### Rule System

Each rule is a `rules.F` struct containing:
1. A `RegExp` pattern to match text
2. An `Applier` function that modifies the `Context` with extracted values

Rules are grouped by language in `rules/{lang}/`:
- `en/` - English
- `ru/` - Russian
- `br/` - Brazilian Portuguese
- `zh/` - Chinese
- `nl/` - Dutch
- `common/` - Language-agnostic patterns (slash dates like "12/25/2023")

### Parsing Flow

1. All rules are matched against input text
2. Matches within `Distance` (default 5 chars) are clustered together
3. Matched rules are applied to a `Context` in order
4. `Context.Time()` combines accumulated values with base time to produce result

### Adding New Rules

Create a function returning `rules.Rule` that uses `rules.F`:
```go
func MyRule(s rules.Strategy) rules.Rule {
    return &rules.F{
        RegExp: regexp.MustCompile(`(?i)pattern`),
        Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
            // m.Captures contains regex capture groups
            // Set values on c (c.Hour, c.Day, c.Duration, etc.)
            return true, nil
        },
    }
}
```

Add the rule to the language's `All` slice in `{lang}/{lang}.go`.