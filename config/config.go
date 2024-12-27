package config

import (
    "os"
    "path/filepath"
)

const ConfigPath = ".config/backup-manager/config.yml"

func GetConfigPath() string {
    homeDir, _ := os.UserHomeDir()
    return filepath.Join(homeDir, ConfigPath)
}

func LoadConfig() {
    // Načtení konfigurace ze souboru.
}

func SaveConfig() {
    // Uložení konfigurace do souboru.
}
