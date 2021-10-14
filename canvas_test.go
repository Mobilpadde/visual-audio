package visual_test

import (
	"testing"

	"github.com/Mobilpadde/visual-audio"
)

func TestWaves(t *testing.T) {
	wavPath := "./example/sample.wav"
	brandPath := "./example/kiosk-branding.png"
	nSamples := 200

	samples, err := visual.ReadPath(wavPath, nSamples)
	if err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}

	if len(samples) != nSamples {
		t.Fatalf("expected samples length to be %d, got: %d", nSamples, len(samples))
	}

	canvas := visual.Blank(samples, 2, 5, 500, false)

	if _, err := canvas.BrandingPath(brandPath, 0.9, false); err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}

	canvas.Waves(228, 71, 54, 20)

	if err := canvas.Save("test-waves.png"); err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}
}

func TestCircleWaves(t *testing.T) {
	wavPath := "./example/sample.wav"
	brandPath := "./example/kiosk-branding.png"
	nSamples := 200

	samples, err := visual.ReadPath(wavPath, nSamples)
	if err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}

	if len(samples) != nSamples {
		t.Fatalf("expected samples length to be %d, got: %d", nSamples, len(samples))
	}

	canvas := visual.Blank(samples, 50, 5, 1000, true)

	if _, err := canvas.BrandingPath(brandPath, 0.9, false); err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}

	canvas.CircleWaves(228, 71, 54, 200)

	if err := canvas.Save("test-circle.png"); err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}
}

func BenchmarkWave(b *testing.B) {
	wavPath := "./example/sample.wav"
	nSamples := 250

	samples, _ := visual.ReadPath(wavPath, nSamples)
	canvas := visual.Blank(samples, 2, 5, 500, false)

	for i := 0; i < b.N; i++ {
		canvas.Waves(228, 71, 54, 20)
	}
}
