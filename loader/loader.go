package main

import (
    "os/exec"
    "log"
    "path/filepath"
)

func runPythonMod(modPath string) {
    script := "scripts/run_mod.py"
    cmd := exec.Command("python3", script, modPath)
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

// Example function to run a specific mod
func RunMod(modName string) {
    modPath := filepath.Join("content/mods", modName, "mod.py")
    runPythonMod(modPath)
}
