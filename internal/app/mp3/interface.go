package mp3

// Interface for Player
type Player interface {
    Play(filePath string) error
    Stop() error
    Pause() error
    Close() error
    IsPlaying() error
}
