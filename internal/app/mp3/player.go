package mp3

import (
    "fmt"
    "github.com/hajimehoshi/go-mp3"
    "github.com/hajimehoshi/oto/v2"
    "os"
    "time"
)

//Mp3player struct
type Mp3Player struct {
    filePath  string
    isPlaying bool
    player    oto.Player
}

// Mp3Player constructor
func New() *Mp3Player {
    return &Mp3Player{
        filePath:  "",
        isPlaying: false,
        player:    nil,
    }
}

// Play mp3 file with filePath
func (pl *Mp3Player) Play(filePath string) error {
    f, err := os.Open(filePath)
    if err != nil {
        return err
    }

    defer f.Close()

    d, err := mp3.NewDecoder(f)
    if err != nil {
        return err
    }

    c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
    if err != nil {
        return err
    }

    <-ready

    p := c.NewPlayer(d)
    pl.player = p

    defer p.Close()
    p.Play()

    pl.isPlaying = p.IsPlaying()

    fmt.Printf("Length %d[bytes]\n", d.Length())
    for {
        time.Sleep(time.Second)
        if !pl.isPlaying {
            break
        }

    }

    return nil
}

// Pause composition
func (pl *Mp3Player) Pause() error {
    pl.player.Pause()

    return nil
}

// Stop composition
func (pl *Mp3Player) Stop() error {
    pl.isPlaying = false

    return nil
}

// Close composition
func (pl *Mp3Player) Close() {
    if pl.player != nil {
        pl.player.Close()
    }
}

// Check is playing composition
func (pl *Mp3Player) IsPlaying() bool {
    return pl.isPlaying
}
