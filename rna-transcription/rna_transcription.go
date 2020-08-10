package strand

func ToRNA(dna string) string {
	var rna = ""
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
