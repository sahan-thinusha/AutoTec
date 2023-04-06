package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"log"
)

func GeneratePreRepairEstimatePDF() {
	fmt.Println("xxx")

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4

	fmt.Println("Sss")
	err := pdf.AddTTFFont("Roboto-Regular", "./res/Roboto-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("Roboto-Regular", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.AddHeader(func() {
		pdf.SetY(5)
		pdf.Cell(nil, "header")
	})
	pdf.AddFooter(func() {
		pdf.SetY(825)
		pdf.Cell(nil, "footer")
	})

	pdf.AddPage()
	pdf.SetY(50)
	pdf.Text("Customer ID :" + "cst")
	pdf.SetX(250)
	pdf.Text("Customer Name :" + "Sahan")
	pdf.SetX(0)
	pdf.SetY(400)
	pdf.Text("page 1 content")
	pdf.AddPage()
	pdf.SetY(400)
	pdf.Text("page 2 content")
	err = pdf.WritePdf("header-footer.tmp.pdf")
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println("zcxc")
}
