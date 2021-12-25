package guard

import (
	"strings"
	"sync"
	"testing"
)

func TestValue(t *testing.T) {
	wg := sync.WaitGroup{}
	g := New[string]("dog")
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			g.Write(func(v string) string {
				return v + "s"
			})
		}()
	}
	wg.Wait()
	g.Read(func(v string) {
		if count := strings.Count(v, "s"); count != 20 {
			t.Logf("count is %v", count)
		}
		t.Logf("v is %v", g.inner)
	})
}
