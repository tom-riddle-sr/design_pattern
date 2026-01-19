package main

type oldMp3Player struct{}

func (p *oldMp3Player) PlayMp3(filename string) string {
	return "Playing mp3 file: " + filename
}

type newMediaPlayer interface {
	Play(filename string, format string) string
}
type mediaAdapter struct {
	oldPlayer *oldMp3Player
}

func newMediaAdapter() *mediaAdapter {
	return &mediaAdapter{oldPlayer: &oldMp3Player{}}
}

func (a *mediaAdapter) Play(filename string, format string) string {
	if format == "mp3" {
		return a.oldPlayer.PlayMp3(filename)
	}
	return "Invalid media. " + format + " format not supported"
}

func main() {
	var player newMediaPlayer = newMediaAdapter()

	println(player.Play("song.mp3", "mp3"))
	println(player.Play("video.mp4", "mp4"))
}
