# Renovate gomodTidyAll Test Repository

This is a test repository to verify Renovate's graph-aware gomodTidyAll functionality.

## Structure

```
.
├── shared/          # Shared library module
├── api/            # API module (depends on shared)
├── sdk/            # SDK module (depends on shared)
└── web/            # Web application (depends on api and sdk)
```

## Dependency Graph

```
shared/go.mod
├── api/go.mod
└── sdk/go.mod
    └── web/go.mod
```

When the shared module is updated, gomodTidyAll should process all modules in dependency order.