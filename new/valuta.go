package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main4() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var bankA []float64
		bankA = make([]float64, 6)

		for j := 0; j < 6; j++ {
			var n, m int
			fmt.Fscan(in, &n, &m)
			bankA[j] = float64(m) / float64(n)
		}

		var bankB []float64
		bankB = make([]float64, 6)
		for j := 0; j < 6; j++ {
			var n, m int
			fmt.Fscan(in, &n, &m)
			bankB[j] = float64(m) / float64(n)
		}

		var bankC []float64
		bankC = make([]float64, 6)
		for j := 0; j < 6; j++ {
			var n, m int
			fmt.Fscan(in, &n, &m)
			bankC[j] = float64(m) / float64(n)
		}

		max := math.Max(bankC[0], math.Max(bankA[0], bankB[0]))

		t1 := bankA[1] * bankB[5]
		t2 := bankB[1] * bankA[5]
		t3 := bankA[1] * bankC[5]
		t4 := bankC[1] * bankA[5]
		t5 := bankB[1] * bankC[5]
		t6 := bankC[1] * bankB[5]

		max = math.Max(max, math.Max(t1, math.Max(t2, math.Max(t3, math.Max(t4, math.Max(t5, t6))))))

		t1 = bankA[0] * bankB[3] * bankC[5]
		t2 = bankA[0] * bankC[3] * bankB[5]
		t3 = bankC[0] * bankA[3] * bankB[5]
		t4 = bankC[0] * bankB[3] * bankA[5]
		t5 = bankB[0] * bankA[3] * bankC[5]
		t6 = bankB[0] * bankC[3] * bankA[5]

		max = math.Max(max, math.Max(t1, math.Max(t2, math.Max(t3, math.Max(t4, math.Max(t5, t6))))))

		t1 = bankA[0] * bankB[2] * bankC[0]
		t2 = bankA[0] * bankC[2] * bankB[0]
		t3 = bankC[0] * bankA[2] * bankB[0]
		t4 = bankC[0] * bankB[2] * bankA[0]
		t5 = bankB[0] * bankA[2] * bankC[0]
		t6 = bankB[0] * bankC[2] * bankA[0]

		max = math.Max(max, math.Max(t1, math.Max(t2, math.Max(t3, math.Max(t4, math.Max(t5, t6))))))

		t1 = bankA[1] * bankB[4] * bankC[0]
		t2 = bankA[1] * bankC[4] * bankB[0]
		t3 = bankC[1] * bankA[4] * bankB[0]
		t4 = bankC[1] * bankB[4] * bankA[0]
		t5 = bankB[1] * bankA[4] * bankC[0]
		t6 = bankB[1] * bankC[4] * bankA[0]

		max = math.Max(max, math.Max(t1, math.Max(t2, math.Max(t3, math.Max(t4, math.Max(t5, t6))))))

		fmt.Println(max)
	}
}
