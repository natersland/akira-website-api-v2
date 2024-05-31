package commonhelpers

import "testing"

func TestGenerateUUID(t *testing.T) {
	t.Run("uuid must be return correct value", func(t *testing.T) {
		// Arrange - Nothing to arrange
		// Act
		result := GenerateUUID()

		// Assert
		if result == "" {
			t.Errorf("expected non-empty string, got empty string")
		}
	},
	)

	t.Run("returns a string of length 36", func(t *testing.T) {
		// Arrange - Nothing to arrange
		// Act
		result := GenerateUUID()

		// Assert
		if len(result) != 36 {
			t.Errorf("expected string of length 36, got string of length %d", len(result))
		}
	},
	)

	t.Run("returns a string in the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", func(t *testing.T) {
		// Arrange - Nothing to arrange
		// Act
		result := GenerateUUID()

		// Assert
		if len(result) != 36 {
			t.Errorf("expected string of length 36, got string of length %d", len(result))
		}
		if result[8] != '-' || result[13] != '-' || result[18] != '-' || result[23] != '-' {
			t.Errorf("expected string in the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx, got %s", result)
		}
	},
	)
}
