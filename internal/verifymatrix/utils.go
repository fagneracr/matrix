package verifymatrix

import "regexp"

/*ArrayisValid - validate sequence input is only: B|U|D|H
  input: array to validade
  output: if is valid or not*/
func ArrayisValid(in []string) (valid bool) {
	for _, s := range in {
		matched, err := regexp.MatchString(`[^B|U|D|H]`, s)
		if err != nil || matched == true {
			return false
		}
	}
	return true

}

func isValid(i int, j int, r int, c int) bool {
	if i < 0 {
		return false
	}
	if i >= r {
		return false
	}
	if j >= c {
		return false
	}
	if j < 0 {
		return false
	}
	return true
}
