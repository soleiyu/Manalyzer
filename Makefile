default:
	go run main.go
	gnuplot plot.txt
	eog res.png
