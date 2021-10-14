package visual

import (
	"io"
	"os"

	"github.com/mjibson/go-dsp/wav"
)

// Read reads a given reader (r) and returns N samples (nSamples)
func Read(r io.Reader, nSamples int) ([]int16, error) {
	w, err := wav.New(r)
	if err != nil {
		return []int16{}, err
	}

	samples, err := w.ReadSamples(nSamples)
	if err != nil {
		return []int16{}, err
	}

	return samples.([]int16), nil
}

// Read reads a given file (path) and returns N samples (nSamples)
func ReadPath(path string, nSamples int) ([]int16, error) {
	f, err := os.Open(path)
	if err != nil {
		return []int16{}, err
	}
	defer f.Close()

	return Read(f, nSamples)
}
