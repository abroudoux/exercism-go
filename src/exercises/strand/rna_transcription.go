package strand

import "strings"

func ToRNA(dna string) string {
	if dna == "" {
		return dna
	}

	var rna strings.Builder
	for i := 0; i < len(dna); i++ {
		switch string(dna[i]) {
		case "G":
			rna.WriteString("C")
		case "C":
			rna.WriteString("G")
		case "T":
			rna.WriteString("A")
		case "A":
			rna.WriteString("U")
		}
	}

	return rna.String()
}
