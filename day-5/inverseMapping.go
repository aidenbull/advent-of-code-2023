package main

func InvertMapping(input []FertilizerMap) []FertilizerMap {
	out := make([]FertilizerMap, 0)
	for _, m := range(input) {
		out = append(out, FertilizerMap{
			m.SrcStart,
			m.DestStart,
			m.Length,
		})
	}
	return out
}