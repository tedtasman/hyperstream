# HYPERSTREAM
A modular system for streaming data from one thing to another. Modularity scope creep got me again...

## Project Layout
Application code lives in modules within `internal/`:

| Module | Purpose |
| ------ | ------- |
| `control` | control-plane logic for scheduling and managing threads | 
| `orchestration` | per job management logic for completing a single poll step |
| `source` | connectors for inbound data sources |
| `transform` | transformation logic for parsing inbound data |
| `sink` | connectors for outbound data sources |
| `contract` | cross-module interfaces