package mp3

type Player interface {
    Play(filePath string) error
    Stop() error
    Pause() error
    Close() error
    IsPlaying() error
}
