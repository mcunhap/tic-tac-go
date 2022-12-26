package app

import "strconv"

func StringToIntSlice(ss []string) ([]int, error) {
	si := make([]int, 0, len(ss))

	for _, s := range ss {
		i, err := strconv.Atoi(s)

		if err != nil {
			return si, err
		}

		si = append(si, i)
	}

	return si, nil
}
