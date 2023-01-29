package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findSRTs() []string {
	files, _ := filepath.Glob("*.srt")
	return files
}

func findVideos() []string {
	result, _ := filepath.Glob("*.mkv")
	mp4s, _ := filepath.Glob("*.mp4")
	result = append(result, mp4s...)
	avis, _ := filepath.Glob("*.avi")
	result = append(result, avis...)
	return result
}

func rename(srt, video string) {
	old := srt
	video_base := strings.TrimSuffix(video, filepath.Ext(video))
	new := video_base + ".srt"
	if old == new {
		fmt.Println("# nothing to do")
		return
	}
	// else
	fmt.Printf("# renaming %q -> %q\n", old, new)
	err := os.Rename(old, new)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	subtitles := findSRTs()
	videos := findVideos()
	if len(subtitles) != 1 || len(videos) != 1 {
		fmt.Println("Error: there should be exactly one video and one subtitle file in the current directory")
		fmt.Println("Supported video formats: .mkv, .mp4, .avi")
		fmt.Println("Supported subtitle format: .srt")
		os.Exit(1)
	}
	// else
	srt := subtitles[0]
	video := videos[0]
	rename(srt, video)
}
