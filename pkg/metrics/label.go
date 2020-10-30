package metrics

// LabelValues ...
type LabelValues []string

// With validates the input, and returns a new aggregate labelValues.
func (lvs LabelValues) With(labelValues ...string) LabelValues {
	if len(labelValues)%2 != 0 {
		labelValues = append(labelValues, "unknown")
	}
	return append(lvs, labelValues...)
}
