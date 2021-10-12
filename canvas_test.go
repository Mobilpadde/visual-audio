package visual_test

import (
	"testing"

	"github.com/Mobilpadde/visual-audio"
)

func TestWave(t *testing.T) {
	wavPath := "./example/sample.wav"
	brandPath := "./example/kiosk-branding.png"
	nSamples := 250

	samples, err := visual.Read(wavPath, nSamples)
	if err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}

	if len(samples) != nSamples {
		t.Fatalf("expected samples length to be %d, got: %d", nSamples, len(samples))
	}

	canvas := visual.Blank(samples, 2, 5, 500)

	if _, err := canvas.Branding(brandPath, 0.9, false); err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}

	canvas.Waves(228, 71, 54, 20)

	if err := canvas.Save("test.png"); err != nil {
		t.Fatalf("expected error to be nil, got: %s", err.Error())
	}
}
