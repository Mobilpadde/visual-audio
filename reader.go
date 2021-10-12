package visual

import (
	"os"

	"github.com/mjibson/go-dsp/wav"
)

// Read reads a given file (path) and returns N samples (nSamples)
func Read(path string, nSamples int) ([]int16, error) {
	f, err := os.Open(path)
	if err != nil {
		return []int16{}, err
	}
	defer f.Close()

	w, err := wav.New(f)
	if err != nil {
		return []int16{}, err
	}

	samples, err := w.ReadSamples(nSamples)
	if err != nil {
		return []int16{}, err
	}

	return samples.([]int16), nil
}
