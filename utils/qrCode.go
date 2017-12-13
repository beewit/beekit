package utils

import (
	"bytes"
	"github.com/beewit/beekit/utils/imgbase64"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
)

func CreateQrCode(url string) (string, error) {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return "", err
	}
	// Scale the barcode to 300*300 pixelscustomer_educational
	qrCode, err = barcode.Scale(qrCode, 300, 300)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	png.Encode(&b, qrCode)

	return imgbase64.FromBuffer(b), nil
}

func CreateQrCodeBytes(url string) (bytes.Buffer, error) {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return bytes.Buffer{}, err
	}
	// Scale the barcode to 300*300 pixelscustomer_educational
	qrCode, err = barcode.Scale(qrCode, 300, 300)
	if err != nil {
		return bytes.Buffer{}, err
	}
	var b bytes.Buffer
	png.Encode(&b, qrCode)

	return b, nil
}
