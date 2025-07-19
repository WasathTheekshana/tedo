package version

// Version information
const (
    Version = "1.0.0"
    AppName = "Tedo"
    Description = "Terminal Todo - A beautiful CLI todo application"
)

// GetVersion returns the current version
func GetVersion() string {
    return Version
}

// GetFullVersion returns version with app name
func GetFullVersion() string {
    return AppName + " v" + Version
}
