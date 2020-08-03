package main

import (
	"./lib"
	"os/exec"
	"fmt"
)

func main() {
	println("hello")

	wav := lib.Wavpars("ss.wav")

//	rd, ld := lib.Datapars16(wav.DATA)
//	ndat := lib.Datamarg16(rd, ld)
//	wav.DATA = ndat

	for i := 0; i < 1200; i++ {
		fmt.Println(i)
		lib.QDatadumpFde(wav, "dump", float32(i) / float32(30),  0.5)
		ofn := fmt.Sprintf("crbox/cr_%05d.png", i)

		cmd := exec.Command("gnuplot", "plot.txt")
		cmd.Start()
		cmd.Wait()

		cmd = exec.Command("mv", "res.png", ofn)
		cmd.Start()
		cmd.Wait()
	}

	lib.Wavout("res.wav", wav)
//	lib.Dft_rn()

}
