# Pricer

## 1. Start the server

```bash
go run ./server
```

Use `--db path/to/pricer.db` to change the default database path.

## 2. Start the client

```bash
go run ./client
```

## * Regenerate gRPC code

```bash
make protoc
```