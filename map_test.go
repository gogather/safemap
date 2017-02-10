package safemap

import "testing"
import "fmt"

func Test_Map(t *testing.T) {
	safeMap := New()
	// m := map[string]interface{}{}
	count := 10000

	go func() {
		for i := 0; i < count; i++ {
			safeMap.Put(fmt.Sprintf("%d", i), i)
			// m[fmt.Sprintf("%d", i)] = i
		}
	}()

	go func() {
		for j := 0; j < count; j++ {
			fmt.Println(safeMap.Get(fmt.Sprintf("%d", j)))
			// fmt.Println(m[fmt.Sprintf("%d", j)])
		}
	}()

	for true {
	}
}

func Test_String(t *testing.T) {
	safeMap := New()
	count := 1000
	for i := 0; i < count; i++ {
		safeMap.Put(fmt.Sprintf("%d", i), i)
	}

	fmt.Println(safeMap.JSON())
}
