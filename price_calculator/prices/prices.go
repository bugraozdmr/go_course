package prices

import (
	"app/conversion"
	"app/iomanager"
	"fmt"
	"strconv"
)

type taxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"` // ignored -
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *taxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return err
		}

		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices
	return nil
}

// direkt buna atanmis bir metod gibi olacaktır
func (job taxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	// _ kabul ediyorum ama kullanmıyorum demek
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)

}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *taxIncludedPriceJob {
	return &taxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
