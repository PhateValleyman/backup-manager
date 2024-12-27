package backup

import (
    "fmt"
)

func Backup(files []string, destination string, local bool) {
    if local {
        fmt.Println("Backing up locally to:", destination)
    } else {
        fmt.Println("Backing up to remote server...")
    }
    // Implementujte logiku zálohování zde.
}
