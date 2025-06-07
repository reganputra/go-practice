package prices

import (
	"fmt"
	"price-calculator/converter"
	"price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLine()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	prices, err := converter.StringToFloats(lines)
	if err != nil {
		fmt.Println("Error converting strings to floats:", err)

		return
	}
	job.InputPrices = prices

}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxedIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedIncludePrice)
	}

	job.TaxIncludedPrices = result
	err := job.IOManager.WriteJson(job)
	if err != nil {
		return
	}
}
