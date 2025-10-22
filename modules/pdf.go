package modules

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/spf13/cobra"
)

func NewPdfCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pdf",
		Short: "Generate a base64-encoded PDF document",
		Run:   runPdf,
	}
}

func runPdf(_ *cobra.Command, _ []string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 16)

	const size = 21 * 1024 * 1024
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteByte(charset[rand.Intn(len(charset))])
	}

	pdf.Cell(40, 10, buffer.String())

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		fmt.Printf("Error generating PDF: %v", err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(buf.Bytes()))
}
