default:
	ffmpeg	-r 30 -i crbox/cr_%05d.png -i ss.wav -vcodec nvenc out.mp4
hoge:
	go run main.go
	gnuplot plot.txt
	eog res.png
