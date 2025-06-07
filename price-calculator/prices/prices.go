package prices

import (
	"fmt"
	"price-calculator/converter"
	"price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
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

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxedIncludePrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedIncludePrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}
