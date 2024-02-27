package location

import "testing"

func TestLocation(t *testing.T) {
	loc := &Location{
		Region: "Santa Catarina",
		City:   "Blumenau",
	}
	t.Run("Equals", func(t *testing.T) {
		other := &Location{
			Region: "Santa Catarina",
			City:   "Blumenau",
		}
		if got := loc.Equals(other); !got {
			t.Errorf("Equals(%v) = %v, want %v", other, got, true)
		}
		other.Region = "Rio Grande do Norte"
		if got := loc.Equals(other); got {
			t.Errorf("Equals(%v) = %v, want %v", other, got, false)
		}
	})
}
