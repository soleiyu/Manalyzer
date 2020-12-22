default:
	ffmpeg	-r 30 -i crbox/cr_%05d.png -i dummy.wav -vcodec nvenc out.mp4
hoge:
	go run main.go
	gnuplot plot.txt
	eog res.png
hoge:
	go run main.go
	gnuplot plot2.txt
	eog res.png
