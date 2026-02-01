package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() string
}

type Playlist struct {
	songs []string
}

func NewPlaylist(songs []string) *Playlist {
	return &Playlist{songs: songs}
}

func (p *Playlist) Iterator() Iterator {
	return &PlaylistIterator{playlist: p}
}

type PlaylistIterator struct {
	playlist *Playlist
	index    int
}

func (it *PlaylistIterator) HasNext() bool {
	return it.index < len(it.playlist.songs)
}

func (it *PlaylistIterator) Next() string {
	if !it.HasNext() {
		return ""
	}
	song := it.playlist.songs[it.index]
	it.index++
	return song
}

func main() {
	fmt.Println("===== Iterator Pattern Example =====")

	playlist := NewPlaylist([]string{"Intro", "Focus", "Break", "Deep Work"})
	it := playlist.Iterator()

	for it.HasNext() {
		fmt.Println("Play:", it.Next())
	}
}
