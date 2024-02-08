package main

func InvertMapping(input MapSet) MapSet {
	out := make([]FertilizerMap, 0)
	for _, m := range(input.Maps) {
		out = append(out, FertilizerMap{
			m.SrcStart,
			m.DestStart,
			m.Length,
		})
	}
	return MapSet{out}
}