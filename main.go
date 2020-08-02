package main

import (
	"./lib"
)

func main() {
	println("hello")

	wav := lib.Wavpars("ss.wav")

//	rd, ld := lib.Datapars16(wav.DATA)
//	ndat := lib.Datamarg16(rd, ld)
//	wav.DATA = ndat

	lib.Datadump16(wav, "dump", 0.1)

	lib.Wavout("res.wav", wav)
//	lib.Dft_rn()

}
