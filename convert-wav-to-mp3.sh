#!/bin/bash
#ffmpeg -i input.wav -vn -ar 44100 -ac 2 -ab 192k -f mp3 output.mp3

ffmpeg -i $1 -vn -ar 44100 -ac 2 -ab 192k -f mp3 $2