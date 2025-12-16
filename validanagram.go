// You can edit this code!
// Click here and start typing.
package main


func isAnagram(s1 string, s2 string) bool {

	if len(s1) != len(s2) {

		return false
	}

	m := make(map[rune]int)

	for _, v := range s1 {
		m[v] = m[v] + 1
	}

	for _, v := range s2 {

		if _, ok := m[v]; !ok {
			return false
		} 
		m[v]--
	}

	for _, value := range m {
		if value != 0 {
			return false
		}

	}

	return true

}
