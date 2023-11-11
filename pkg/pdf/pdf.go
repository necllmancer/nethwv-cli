package pdf

import (
	"fmt"
	"nethwv-cli/pkg/filefetcher"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(fileUrls []string, output string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 12)

	for _, url := range fileUrls {
		content, err := filefetcher.FetchFileContent(url)
		if err != nil {
			fmt.Println("GeneratePDF error: ", err)
			return err
		}
		if content != "" {
			pdf.AddPage()
			pdf.SetXY(10, 10) // 設定テキスト開始位置
			pdf.MultiCell(0, 10, url+"\n"+content, "", "", false)
		}
	}

	err := pdf.OutputFileAndClose(output)
	if err != nil {
		fmt.Println("pdf.OutputFileAndClose(output)", err)
		return err
	}

	return nil
}
