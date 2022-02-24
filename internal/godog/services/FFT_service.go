package services

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"os"
	"zarbat_test/internal/logging"
	l "zarbat_test/internal/logging"

	"github.com/xigh/go-wavreader"
)

const (
	LENGTH = 4000
)

func GetFrequencies(filePath string, expectedFrequency int, similarity int) (err error) {

	r, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	wr, err := wavreader.New(r)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s: %dHz, %d channels, %d samples, %v\n",filePath, wr.Rate(), wr.Chans(), wr.Len(), wr.Duration())

	samples := make([]float64, LENGTH)
	for i := uint64(0); i < LENGTH; i++ {
		s, err := wr.At(0, i)
		if err != nil {
			log.Fatal(err)
		}
		samples[i] = float64(s)
	}

	freqs := fft(samples)
	var max float64 = 0
	var maxi int = 0
	for i := range freqs {
		m := cmplx.Abs(freqs[i])
		if m > max {
			max = m
			maxi = i
		}
	}
	fmt.Printf("Maximum: %f", max)
	l.Debug.Printf("Maximum: %f", max)
	var fq []int
	for i := range freqs {
		m := cmplx.Abs(freqs[i])
		if m > max/2 {
			logging.Debug.Println(i, m)
			fq = append(fq, int(wr.Rate())/LENGTH*i)
		}
	}
	fmt.Printf("The maximum index is %d\n", maxi)
	l.Debug.Printf("The maximum index is %d\n", maxi)
	for j := range fq {
		if float64(fq[j]) < float64(expectedFrequency)*float64(1.10) && float64(fq[j]) > float64(expectedFrequency)*float64(0.9) {
			return nil
		}
	}
	return fmt.Errorf("Error in getting the frequency")
}

func hfft(samples []float64, freqs []complex128, n, step int) {
	if n == 1 {
		freqs[0] = complex(samples[0], 0)
		return
	}

	half := n / 2

	hfft(samples, freqs, half, 2*step)
	hfft(samples[step:], freqs[half:], half, 2*step)

	for k := 0; k < half; k++ {
		a := -2 * math.Pi * float64(k) / float64(n)
		e := cmplx.Rect(1, a) * freqs[k+half]

		freqs[k], freqs[k+half] = freqs[k]+e, freqs[k]-e
	}
}

func fft(samples []float64) []complex128 {
	n := len(samples)
	freqs := make([]complex128, n)
	hfft(samples, freqs, n, 1)
	return freqs
}
