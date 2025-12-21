package commons

type TimeWindow struct {
	Start uint64
	End   uint64
}

const MinTimeWindow = 3 * 24 * 60
