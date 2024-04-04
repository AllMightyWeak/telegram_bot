package generator

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(link string) {
	err := qrcode.WriteFile(link, qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("QR code doesn't gemerate")
	}
}
