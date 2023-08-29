# Install GCC
https://jmeubank.github.io/tdm-gcc/download/

# Fyne doc
https://developer.fyne.io/started/

# Run
go run ./cmd

# Build for windows
fyne package -os windows -icon icon.png -src ./cmd -name mygoapp

# Install playwright
https://github.com/playwright-community/playwright-go

# Test
go test ./...

coverage:
go test ./... -cover
