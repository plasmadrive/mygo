package qsort

import (
	"fmt"
)

func QSort(nn []int) error {
	return QuickSort(nn, 0, len(nn)-1)

}

func QuickSort(nn []int, lo int, hi int) error {
	if lo < 0 {
		return fmt.Errorf("Lo value : %d cannot be less than zero", lo)
	}
	if hi >= len(nn) {
		return fmt.Errorf("Hi value : %d cannot be greater than or equal to %d", hi, len(nn))
	}

	var (
		err error
		p   int
	)
	
	if lo < hi {

		p, err = Partition(nn, lo, hi)
	

		err = QuickSort(nn, lo, p-1)
	
		err = QuickSort(nn, p+1, hi)

	}
	return err
}

func Partition(nn []int, lo int, hi int) (int, error) {

	if lo < 0 {
		return 0, fmt.Errorf("Lo : %d is less than zero", lo)
	}
	if hi >= len(nn) {
		return 0, fmt.Errorf("Hi : %d exceeds slice length : %d", hi, len(nn))
	}

	if lo > hi {
		return 0, fmt.Errorf("Lo : %d is greater than Hi : %d", lo, hi)
	}

	pivot := nn[hi]

	i := lo
	for j := lo; j <= hi-1; j++ {
		if nn[j] <= pivot {
			swap(nn, i, j)
			i += 1
			_ = "breakpoint"
		}

	}
	swap(nn, i, hi)
	_ = "breakpoint"
	return i, nil
}

func swap(nn []int, i int, j int) error {
	if i < 0 {
		return fmt.Errorf("Index %d is less than zero", i)
	}

	if j < 0 {
		return fmt.Errorf("Index %d is less than zero", j)
	}

	if i > len(nn)-1 {
		return fmt.Errorf("INdex %d is greater than %d", i, len(nn)-1)
	}

	if j > len(nn)-1 {
		return fmt.Errorf("INdex %d is greater than %d", j, len(nn)-1)
	}

	t := nn[i]
	nn[i] = nn[j]
	nn[j] = t
	return nil
}
