package qsort

import (
	"testing"
	"math/rand"
)

func TestQSort(t *testing.T) {
	nn := []int{1, 3, 2, 5, 4}
	err := QSort(nn)
	if err != nil {
		t.Fatalf("Quicksort returned an error", err)
	}
	sorted := []int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		if nn[i] != sorted[i] {
			t.Fatalf("Slice is is not sorted at postition %d. Actual value %d, expected valued %d", i, nn[i], sorted[i])
		}
	}

}

func TestQuickSortParameters(t *testing.T) {
	// test lo < 0
	nn := []int {3,2,1}
	var (
		err error
	)
	err = QuickSort(nn,-1,2)
	if err == nil {
		t.Errorf("QuickSort should have returned an error for lo value -1")
	}

	err = QuickSort(nn,0,3)
	if err == nil {
		t.Errorf("Quicksort shouldh have returned an error for hi value 3")
	}

	

	
}

func TestPartition(t *testing.T) {

	nn := []int{7, 2, 1, 8, 6, 3, 5, 4}
	part := []int{2, 1, 3, 4, 6, 7, 5, 8}
	pi := 3
	p, err := Partition(nn, 0, len(nn)-1)
	if err != nil {
		t.Fatalf("Partition returned error", err)
	}
	for i := 0; i < 8; i++ {

		if nn[i] != part[i] {
			t.Fatalf("Error crateing partition. Expected value at index %d : %d. Actual Value %d", i, part[i], nn[i])
		}
		if p != pi {
			t.Fatalf("Partition index incorrect: Expected %d, Actual %d", pi, p)
		}
	}

}

func BenchmarkPartition(b *testing.B) {
	nn := generateSlice(100)
	b.ResetTimer()
	for i :=0; i< b.N; i++ {
		Partition(nn,0,99)
	}
}

func generateSlice(len int) []int {
	const maxelementsize int = 1000
	var nn []int
	if len > 0 {
		rand.Seed(1)
		nn = make([]int,len)
		for i :=0;i<len;i++ {
			nn[i] = rand.Intn(maxelementsize)
		}
	}
	return nn
}

func TestPartitionLoGreaterThanHi(t *testing.T) {
	nn := []int{1, 2, 3}
	_, err := Partition(nn, 2, 1)
	if err == nil {
		t.Fatalf("Partition should have returned an error since Lo is greater than Hi")
	}
}

func TestHiGreaterThanLength(t *testing.T) {
	nn := []int{1, 2, 3}
	_, err := Partition(nn, 0, 3)
	if err == nil {
		t.Fatalf("Partition should have returned an error since HI is greater than slice length")
	}
}

func TestLoLessThanZero(t *testing.T) {
	nn := []int{1, 2, 3}
	_, err := Partition(nn, -1, 2)
	if err == nil {
		t.Fatalf("Partition should have returned an error since Lo less than zero")
	}
}

func TestSwap(t *testing.T) {
	var (
		nn []int
	)
	nn = []int{1, 2, 3, 4}

	swap(nn, 0, 3)
	if nn[0] != 4 {
		t.Errorf("Swap Expected %d, Actual %d", 4, nn[0])
	}

	if nn[3] != 1 {
		t.Errorf("Swap Expected %d, Actual %d at index %d", 1, nn[3], 3)
	}

	swap(nn, 1, 2)
	if nn[1] != 3 {
		t.Errorf("Swap Expected %d, Actual %d at index %d", 3, nn[1], 1)
	}

	if nn[2] != 2 {
		t.Errorf("Swap Expected %d, Actual %d at index %d", 2, nn[2], 2)
	}
}

func TestSwapIndexLessThanZero(t *testing.T) {
	nn := []int{1, 2, 3, 4}
	err := swap(nn, -1, 0)
	if err == nil {
		t.Fatalf("Swap should have returned an error since -1 less than zero")
	}
}

func TestSwapIndexGreaterThanSliceLength(t *testing.T) {
	nn := []int{1, 2, 3, 4}
	err := swap(nn, 0, 4)
	if err == nil {
		t.Fatalf("Swap should have returned an error since 4 > %d", len(nn)-1)
	}
}
