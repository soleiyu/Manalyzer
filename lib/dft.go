package lib

import (
	"math"
	"strconv"
	"fmt"
	"os"
)

func DumpDftmde(inp Wav, fn string, tbeg, tlim float32) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	bof := int(float32(inp._SRate) * tbeg)
	num := int(float32(inp._SRate) * tlim)
	rd, _ := Dataparsf64(inp.DATA)

	if len(rd) < bof + 2205 {
		bof = len(rd) - 2205
	}

	for i := 0; i < num / 4; i++ {
		bofs := bof + i * 4
		file.WriteString(strconv.Itoa(i))
		file.WriteString(" ")
		sfv := fmt.Sprintf("%f", Dft_mix(rd[bofs: bofs + 2204], 1) )
		file.WriteString(sfv)
		file.WriteString(" ")
		sfv = fmt.Sprintf("%f", Dft_mix(rd[bofs: bofs + 440], 1) )
		file.WriteString(sfv)
		file.WriteString("\n")
	}
}


func DumpDftm(inp Wav, fn string, tbeg float32) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	bof := int(float32(inp._SRate) * tbeg)
	rd, _ := Dataparsf64(inp.DATA)

	if len(rd) < bof + 2205 {
		bof = len(rd) - 2205
	}

	for i := 2; i < 2205; i ++ {
		sfv := fmt.Sprintf("%d %f\n", 44100 / i, Dft_mix(rd[bof: bof + i], 1) )
		file.WriteString(sfv)
	}
}

func Dft_mix(darr []float64, k float64) float64 {
	cv := 2.0 * k * math.Pi / float64(len(darr))
	res := 0.0

	for i := 0; i < len(darr); i++ {
		res += darr[i] * math.Cos((float64(i) * cv))
		res += darr[i] * math.Sin((float64(i) * cv))
	}

	return res / float64(len(darr))
}


func Dft_rn(darr []float64, k float64) float64 {
	cv := 2.0 * k * math.Pi / float64(len(darr))
	res := 0.0

	for i := 0; i < len(darr); i++ {
		res += darr[i] * math.Cos((float64(i) * cv))
	}

	return res / float64(len(darr))
}

func Dft_im(darr []float64, k float64) float64 {
	cv := 2.0 * k * math.Pi / float64(len(darr))
	res := 0.0

	for i := 0; i < len(darr); i++ {
		res += darr[i] * math.Sin((float64(i) * cv))
	}

	return res / float64(len(darr))
}

