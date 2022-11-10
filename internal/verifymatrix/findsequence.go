package verifymatrix

import "strings"

/*FindSequence - verify if exists sequence
  input: []string with sequences, from API
  output: quantity of sequence found*/
func FindSequence(in []string) (countFound int) {
	toFind := []string{"BBBB", "UUUU", "DDDD", "HHHH"}
	for _, x := range toFind {
		for _, a := range in {
			if strings.Contains(a, x) {
				countFound++
			}
		}
	}
	return countFound

}
