object dft
{
	def main(args: Array[String])
	{

		var pnum = 64

		var cvarr: Array[Double] = Array(1, 4, 8, 5, 3.5, 4.5, 7, -2, 3 , -5, -1, -7)

//		var carr = wavgen.cosPlot(pnum, cvarr) 
//		var carr = cosPlot(pnum, 1, 1) 
		var carr = wavgen.cramPlot(pnum, 3, 10) 

		var drnarr: Array[Double] = Array.empty
		(0 until pnum).foreach (i =>
			drnarr :+= dft_rn(carr, (i + 1.0) / 2.0))

		var dimarr: Array[Double] = Array.empty
		(0 until pnum).foreach (i =>
			dimarr :+= dft_im(carr, (i + 1.0) / 2.0))

		var ncarr = wavgen.hcosPlot(pnum, drnarr)
		var ncari = wavgen.hsinPlot(pnum, dimarr)

		for (i <- 0 until pnum) 
		{
			printf("%d ", i + 1)
			printf("%f ", carr(i))
			printf("%f ", dft_rn(carr, (i + 1.0) / 2.0))
			printf("%f ", dft_im(carr, (i + 1.0) / 2.0))
			println(ncarr(i) + ncari(i))
		}
	}

	def dft_rn(arr: Array[Double], k: Double): Double =
	{
		var cv = 2.0 * k * math.Pi / arr.length 
		var res = 0.0

		(0 until arr.length).foreach (i =>
			res += arr(i) * math.cos(i * cv)
		)

		res / arr.length
	}

	def dft_im(arr: Array[Double], k: Double): Double =
	{
		var cv = 2.0 * k * math.Pi / arr.length 
		var res = 0.0

		(0 until arr.length).foreach (i =>
			res += arr(i) * math.sin(i * cv)
		)

		res / arr.length
	}
}

