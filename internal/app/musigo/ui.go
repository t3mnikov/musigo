package musigo

import (
    "fmt"
    "github.com/gotk3/gotk3/gdk"
    "github.com/gotk3/gotk3/gtk"
    "github.com/t3mnikov/musigo/internal/app/mp3"
    "log"
)

// A user interface
type UI struct {
    win       *gtk.ApplicationWindow
    listBox   *gtk.ListBox
    mp3Player *mp3.Mp3Player //use interface here
}

// UI constructor
func NewUI(app *gtk.Application, mp3pl *mp3.Mp3Player) (*UI, error) {
    ui := &UI{}

    ui.mp3Player = mp3pl
    err := ui.initWindow(app)
    if err != nil {
        return nil, err
    }
    return ui, nil
}

// Add item to list
func (ui *UI) AddToList(names ...string) {
    for _, name := range names {
        l, _ := gtk.LabelNew(name)
        l.SetHAlign(gtk.ALIGN_START)

        lr, _ := gtk.ListBoxRowNew()
        lr.Add(l)

        ui.listBox.Insert(lr, -1)
        ui.listBox.ShowAll()
    }
}

// Initialize a main window with UI
func (ui *UI) initWindow(app *gtk.Application) error {
    win, err := gtk.ApplicationWindowNew(app)
    if err != nil {
        //log.Fatal("Unable to create window:", err)
        return err
    }

    ui.win = win

    win.SetTitle("Musigo player")
    win.Connect("destroy", func() {
        //when the window closes, we disconnect from gtk
        gtk.MainQuit()
    })

    ui.initInterface()

    win.SetDefaultSize(500, 600)
    win.ShowAll()
    gtk.Main()

    return nil
}

// Init user interface
func (ui *UI) initInterface() {
    buttonPlay, _ := gtk.ButtonNew()
    buttonPlay.SetLabel("Play")

    buttonPause, _ := gtk.ButtonNew()
    buttonPause.SetLabel("Pause")

    buttonStop, _ := gtk.ButtonNew()
    buttonStop.SetLabel("Stop")

    buttonOpen, _ := gtk.ButtonNew()
    buttonOpen.SetLabel("Open...")

    buttonsGrid, _ := gtk.GridNew()
    buttonsGrid.SetRowSpacing(1)
    buttonsGrid.SetColumnSpacing(5)
    buttonsGrid.SetColumnHomogeneous(true)
    buttonsGrid.Add(buttonPlay)
    buttonsGrid.Add(buttonPause)
    buttonsGrid.Add(buttonStop)
    buttonsGrid.Add(buttonOpen)

    listBox, _ := gtk.ListBoxNew()
    ui.listBox = listBox

    listBox.Connect("key-release-event", func(w *gtk.ListBox, e *gdk.Event) {
        fmt.Println("released")

        eventKey := gdk.EventKey{e}
        if eventKey.KeyVal() == gdk.KEY_Delete {
            r := w.GetSelectedRow()
            listBox.Remove(r)
        }

    })

    listBox.Connect("button-press-event", func(w *gtk.ListBox, e *gdk.Event) {
        event := gdk.EventButton{e}
        if event.Type() == gdk.EVENT_2BUTTON_PRESS {
            fmt.Println("double clicked")
            filePath := ui.getCurrentFilePath()

            err := ui.mp3Player.Play(filePath)
            if err != nil {
                log.Fatal(err)
            }
        }

    })

    statusBar, _ := gtk.StatusbarNew()

    box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
    box.Add(buttonsGrid)
    box.Add(listBox)
    box.Add(statusBar)

    dialogOpenFile, _ := gtk.FileChooserDialogNewWith1Button("Open audio files", ui.win, gtk.FILE_CHOOSER_ACTION_OPEN, "_Open", gtk.RESPONSE_ACCEPT)
    ff, _ := gtk.FileFilterNew()
    ff.AddPattern("*.mp3")
    dialogOpenFile.SetFilter(ff)
    dialogOpenFile.Connect("response", func() {
        dialogOpenFile.Hide()
        fileName := dialogOpenFile.GetFilename()
        ui.AddToList(fileName)
        fmt.Println(fileName)
    })

    buttonPlay.Connect("clicked", func() {
        fmt.Println("played")

        if ui.mp3Player.IsPlaying() {
            ui.mp3Player.Stop()
        }

        filePath := ui.getCurrentFilePath()

        go ui.mp3Player.Play(filePath)

        fmt.Println(filePath)
    })

    buttonStop.Connect("clicked", func() {
        fmt.Println("stopped")

        if ui.mp3Player.IsPlaying() {
            ui.mp3Player.Stop()
        }
    })

    buttonPause.Connect("clicked", func() {
        fmt.Println("paused")

        if ui.mp3Player.IsPlaying() {
            ui.mp3Player.Pause()
        }
    })

    buttonOpen.Connect("clicked", func() {
        //fmt.Println("clicked")
        dialogOpenFile.ShowAll()
    })

    ui.win.Add(box)
}

// Get current filePath for playing composition
func (ui *UI) getCurrentFilePath() string {
    listBoxRow := ui.listBox.GetSelectedRow()
    lw, err := listBoxRow.GetChild() // Label
    if err != nil {
        log.Fatal(err)
    }

    pl := lw.ToWidget()
    l := *pl
    filePath, _ := l.GetProperty("label")

    return filePath.(string)
}
