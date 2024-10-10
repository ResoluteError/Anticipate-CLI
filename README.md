# Setup from scratch for your own projects

- `mkdir anticipate`
- `cd anticipate`
- `go mod init anticipate`
- `go get github.com/spf13/cobra`
- `go get github.com/spf13/viper`
- `go get github.com/manifoldco/promptui`
- `goreleaser init`

# Local Testing

- `go run main.go add 2025-03-17 -d "Birthday"`
- `go run main.go countdown`
- `go run main.go remove`

# Deploying

- Create and push a branch called `release/[VERSION]` (replace version with something like 0.0.1)
