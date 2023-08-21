package musigo

import (
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
    "github.com/t3mnikov/musigo/internal/app/mp3"
    "log"
    "os"
)

const (
    appID = "org.gtk.musigo"
)

// Run player
func RunPlayer() {
    mp3pl := mp3.New()

    app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
    if err != nil {
        log.Fatal("Could not create application.", err)
    }

    app.Connect("activate", func() { onActivate(app, mp3pl) })
    app.Connect("deactivate", func() { onDeactivate(app, mp3pl) })

    os.Exit(app.Run(os.Args))
}

// Event on activate
func onActivate(app *gtk.Application, mp3pl *mp3.Mp3Player) {
    _, err := NewUI(app, mp3pl)
    if err != nil {
        log.Fatal("Unable to create window: ", err)
    }
}

// Event on deactivate
func onDeactivate(app *gtk.Application, mp3pl *mp3.Mp3Player) {
    mp3pl.Close()
}
