# golinters

Custom Go linters, usable via `go vet -vettool` or as [golangci-lint module plugins](https://golangci-lint.run/docs/plugins/module-plugins/).

## Linters

| Name | Description | Load Mode |
|---|---|---|
| `println` | Reports `println` builtin usage | syntax |
| `printf` | Reports `fmt.(Fp\|P)rintf*` usage (Fprint only flags `os.Stdout`) | syntax |
| `uselessf` | Reports `*f` calls with no format args (e.g. `Errorf("msg")` → `errors.New`) | syntax |
| `deferfor` | Reports `defer` directly inside a `for` loop | syntax |
| `errresp` | Reports `ErrorResponse()` not immediately followed by `return` | syntax |
| `mustcheck` | Reports `.MustVerb()` zero-arg calls outside `init` and package-level declarations | syntax |
| `fatal` | Reports `Logger.Fatal*` calls | types |

## Usage with go vet

Each linter has a standalone binary under `cmd/`:

```bash
go build -o println-checker ./cmd/println
go vet -vettool=./println-checker ./...
```

Or use the combined binary for all linters at once:

```bash
go build -o combined-checker ./cmd/combined
go vet -vettool=./combined-checker ./...
```

## Usage with golangci-lint (module plugin)

Requires [golangci-lint](https://golangci-lint.run/) v2. See [`testproject/`](testproject/) for a working example.

### 1. Add `.custom-gcl.yml`

This tells golangci-lint to build a custom binary with the plugin:

```yaml
version: v2.10.1
plugins:
  - module: 'github.com/edaniels/golinters'
    version: v0.2.0  # or use path: ../golinters for local dev
```

### 2. Add linters to `.golangci.yml`

Enable whichever linters you want:

```yaml
version: "2"

linters:
  enable:
    - println
    - printf
    - uselessf
    - deferfor
    - errresp
    - mustcheck
    - fatal
  settings:
    custom:
      println:
        type: module
      printf:
        type: module
      uselessf:
        type: module
      deferfor:
        type: module
      errresp:
        type: module
      mustcheck:
        type: module
      fatal:
        type: module
```

### 3. Build and run

```bash
golangci-lint custom          # builds ./custom-gcl
./custom-gcl run ./...        # run like normal golangci-lint
```
