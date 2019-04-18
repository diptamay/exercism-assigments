//Package strand contains function which given a DNA strand, return its RNA complement (per RNA transcription).
//
//Both DNA and RNA strands are a sequence of nucleotides.
//The four nucleotides found in DNA are adenine (A), cytosine (C), guanine (G) and thymine (T).
//The four nucleotides found in RNA are adenine (A), cytosine (C), guanine (G) and uracil (U).
//
//Given a DNA strand, its transcribed RNA strand is formed by replacing each nucleotide with its complement:
//G -> C
//C -> G
//T -> A
//A -> U
package strand

//ToRNA given a DNA strand, return its RNA complement (per RNA transcription).
func ToRNA(dna string) string {
	var dnaRna = map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	var ret []rune

	for _, x := range dna {
		rna, ok := dnaRna[x]
		if ok {
			ret = append(ret, rna)
		}
	}
	return string(ret)
}
