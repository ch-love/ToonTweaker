package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/theme"
    "gopkg.in/yaml.v2"
    "os"
    "io/ioutil"
    "net/http"
    "log"
    "strings"
)

// ModConfig represents the mod configuration file structure (from yml)
type ModConfig struct {
    Repo string `yaml:"repo"`
}

// MakeModsPage creates the "Mods" tab layout
func MakeModsPage() fyne.CanvasObject {
    modsFolder := "./content/mods/"
    modFolders, err := ioutil.ReadDir(modsFolder)
    if err != nil {
        log.Fatal(err)
    }

    modList := container.NewVBox()
    
    // Loop through all mod folders and display them
    for _, modFolder := range modFolders {
        if modFolder.IsDir() {
            modConfigPath := modsFolder + modFolder.Name() + "/manege.toontweaker.yml"
            modConfig := loadModConfig(modConfigPath)
            
            // Create mod box layout
            modBox := makeModBox(modFolder.Name(), modConfig)
            modList.Add(modBox)
        }
    }

    return container.NewVScroll(modList)
}

// loadModConfig loads and parses the YAML file from each mod folder
func loadModConfig(path string) ModConfig {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        log.Println("Error reading config:", err)
        return ModConfig{}
    }

    var config ModConfig
    err = yaml.Unmarshal(file, &config)
    if err != nil {
        log.Println("Error parsing YAML:", err)
    }

    return config
}

// makeModBox creates a UI box for each mod, showing the header, gear, and download icons, and toggle
func makeModBox(modName string, config ModConfig) fyne.CanvasObject {
    // Header
    header := widget.NewLabel(modName)
    
    // Paragraph (description placeholder)
    paragraph := widget.NewLabel("Description for " + modName)
    
    // Gear icon (opaque or transparent based on mod having settings)
    gearIcon := canvas.NewImageFromResource(theme.SettingsIcon())
    if !hasSettings(modName) {
        gearIcon.Translucency = 0.5
    }
    
    // Download icon (only visible if updates are available)
    updateAvailable := checkForUpdates(config.Repo)
    downloadIcon := canvas.NewImageFromResource(theme.ContentAddIcon())
    downloadIcon.Hide() // Hide initially

    if updateAvailable {
        downloadIcon.Show() // Show if update available
        // Clicking the icon applies the update
        downloadIcon.OnTapped = func() {
            applyUpdate(config.Repo)
        }
    }
    
    // Toggle switch to enable/disable the mod
    modToggle := widget.NewCheck("Enable Mod", func(checked bool) {
        toggleMod(modName, checked)
    })

    // Layout the mod box with Header, Gear, Download icons, and toggle switch
    modBox := container.NewVBox(
        container.NewBorder(nil, nil, header, container.NewHBox(gearIcon, downloadIcon)),
        paragraph,
        modToggle,
    )

    return modBox
}

// hasSettings checks if the mod has a settings file or options to manage
func hasSettings(modName string) bool {
    settingsPath := "./content/mods/" + modName + "/settings.yml"
    _, err := os.Stat(settingsPath)
    return !os.IsNotExist(err) // Returns true if settings file exists
}

// checkForUpdates checks if the mod repo has updates available on GitHub
func checkForUpdates(repo string) bool {
    if repo == "" {
        return false
    }
    
    // Simulate check for updates (you can expand this with GitHub API)
    resp, err := http.Get("https://api.github.com/repos/" + repo + "/commits")
    if err != nil {
        log.Println("Error checking updates:", err)
        return false
    }
    defer resp.Body.Close()

    return resp.StatusCode == 200 // Assuming 200 means there's new commits
}

// applyUpdate applies the updates from the GitHub repo to the mod
func applyUpdate(repo string) {
    if repo == "" {
        log.Println("No repository specified for update")
        return
    }

    // Placeholder: Actual logic for downloading and applying updates
    log.Println("Applying update for repo:", repo)
}

// toggleMod enables or disables the mod based on user input
func toggleMod(modName string, enabled bool) {
    if enabled {
        log.Println("Mod", modName, "enabled")
    } else {
        log.Println("Mod", modName, "disabled")
    }
}
