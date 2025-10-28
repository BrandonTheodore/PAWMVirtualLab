# Check if Go is installed
if (!(Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Error "Go is not installed. Please install Go first."
    exit 1
}

Write-Host "Running go mod tidy..."
go mod tidy
if ($LASTEXITCODE -ne 0) {
    Write-Error "Error running go mod tidy"
    exit 1
}

Write-Host "Building the application..."
go build -o pawm-virtual-lab.exe
if ($LASTEXITCODE -ne 0) {
    Write-Error "Error building the application"
    exit 1
}

Write-Host "Starting the application..."
.\pawm-virtual-lab.exe