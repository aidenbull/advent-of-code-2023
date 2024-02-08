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

func InvertMultipleMapSets(input []MapSet) []MapSet {
	//Reverse Input
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	
	out := make([]MapSet, 0)
	for _, mapSet := range(input) {
		out = append(out, InvertMapping(mapSet))
	}
	return out
}