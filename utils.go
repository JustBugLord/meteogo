package meteogo

import (
	"strings"
)

func FromUnits[T Stringer](source ...T) string {
	var builder strings.Builder
	for _, unit := range source {
		builder.WriteString(unit.String())
		builder.WriteByte(',')
	}
	return builder.String()
}
