package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"log"

	"github.com/Mobilpadde/visual-audio"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Post("/", clean)
	app.Post("/branding", branding(false))
	app.Post("/repeat", branding(true))

	log.Fatal(app.Listen(":3000"))
}

func clean(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	file := form.File["sample"][0]
	audio, err := file.Open()
	if err != nil {
		return err
	}

	nSamples := 250
	samples, err := visual.Read(audio, nSamples)
	if err != nil {
		return err
	}

	canvas := visual.Blank(samples, 2, 5, 500)
	canvas.Waves(228, 71, 54, 20)

	buf := new(bytes.Buffer)
	im := canvas.Image()
	if err := png.Encode(buf, im); err != nil {
		return err
	}
	r := bytes.NewReader(buf.Bytes())

	c.Response().Header.Set("Content-Type", "image/png")
	c.Status(fiber.StatusCreated)
	io.Copy(c.Response().BodyWriter(), r)
	return nil
}

func branding(repeat bool) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		file := form.File["sample"][0]
		audio, err := file.Open()
		if err != nil {
			return err
		}

		file = form.File["branding"][0]
		brandingImage := fmt.Sprintf("/tmp/%s", file.Filename)
		if err := c.SaveFile(file, brandingImage); err != nil {
			return err
		}

		nSamples := 250
		samples, err := visual.Read(audio, nSamples)
		if err != nil {
			return err
		}

		canvas := visual.Blank(samples, 2, 5, 500)
		if _, err := canvas.BrandingPath(brandingImage, 0.9, repeat); err != nil {
			return err
		}

		canvas.Waves(228, 71, 54, 20)

		buf := new(bytes.Buffer)
		im := canvas.Image()
		if err := png.Encode(buf, im); err != nil {
			return err
		}
		r := bytes.NewReader(buf.Bytes())

		c.Response().Header.Set("Content-Type", "image/png")
		c.Status(fiber.StatusCreated)
		io.Copy(c.Response().BodyWriter(), r)
		return nil
	}
}
