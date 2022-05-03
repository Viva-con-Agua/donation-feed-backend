package dao

// AggregatedTotalMoney is returned when aggregating the total amount of donated money in mongo
type AggregatedTotalMoney struct {
	Currency            string `bson:"_id"`
	TotalDonationAmount int64  `bson:"totalDonationAmount"`
}

func AggregatedTotalMoneyToMap(agg []AggregatedTotalMoney) map[string]int64 {
	result := make(map[string]int64)
	for _, iAgg := range agg {
		result[iAgg.Currency] = iAgg.TotalDonationAmount
	}
	return result
}
