# mvsrt

A simple little program that automatically renames
a subtitle file (`.srt`) to match the name of your video file.

## Motivation

When you download a movie and its subtitle, it happens
quite often that these files have different names, for instance:

    Dexter.S07E04.HDTV.x264-ASAP.mp4
    Dexter - 07x04 - Run.ASAP.English.C.orig.Addic7ed.com.srt

If you want your video player to add the subtitle automatically,
you need to rename the subtitle file to match the name of
the video file:

    Dexter.S07E04.HDTV.x264-ASAP.mp4
    Dexter.S07E04.HDTV.x264-ASAP.srt

However, this repetitive task could be automated, right?

## Supported Formats

Supported video formats: `.mkv`, `.mp4`, `.avi`

Supported subtitle format: `.srt`

## Installation

    $ go install github.com/jabbalaci/mvsrt@latest

It'll give you a binary executable called `mvsrt`.

## Usage

If the executable is in your PATH, then just launch `mvsrt`
in a directory. It'll look for the subtitle file and the video file
and it'll perform the renaming.
