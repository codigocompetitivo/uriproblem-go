package _890
import 	"testing"

func TestTuks(t *testing.T) {
	assertCorrectPlates := func(t testing.TB, expected, output int) {
		t.Helper()
		if expected != output {
			t.Errorf("expected '%d' but got '%d'", expected, output)
		}
	}

	t.Run("to zero consonants and six digits", func(t *testing.T) {
		output := plate(0,6)
		expected := 1000000
		assertCorrectPlates(t,expected, output)
	})

	t.Run("to zero consonants and six digits", func(t *testing.T) {
		output := plate(0,6)
		expected := 1000000
		assertCorrectPlates(t,expected, output)
	})

	t.Run("to zero consonants and six digits", func(t *testing.T) {
		output := plate(2,4)
		expected := 6760000
		assertCorrectPlates(t,expected, output)
	})

	t.Run("to zero consonants and six digits", func(t *testing.T) {
		output := plate(0,0)
		expected := 0
		assertCorrectPlates(t,expected, output)
	})
}