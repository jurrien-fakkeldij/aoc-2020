package transformer

import "strconv"

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func SliceItoa(sa []int) []string {
	si := make([]string, 0, len(sa))
	for _, a := range sa {
		si = append(si, strconv.Itoa(a))
	}
	return si
}
