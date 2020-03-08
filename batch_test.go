package batch

import "testing"

func TestBatch(t *testing.T) {
	type r struct {
		start, end int
	}

	tests := []r{
		r{start: 0, end: 9},
		r{start: 10, end: 19},
		r{start: 20, end: 29},
		r{start: 30, end: 39},
		r{start: 40, end: 49},
		r{start: 50, end: 59},
		r{start: 60, end: 69},
		r{start: 70, end: 79},
		r{start: 80, end: 89},
		r{start: 90, end: 99},
	}

	var ranges []r
	err := Batch(100, 10, func(start, end int) error {
		ranges = append(ranges, r{start, end})
		return nil
	})
	if err != nil {
		t.Error(err)
	}
	if len(ranges) != 10 {
		t.Errorf("len(%q) = %d; want %d", ranges, len(ranges), 10)
	}
	for i, tt := range tests {
		if ranges[i].start != tt.start {
			t.Errorf("start = %d; want %d", ranges[i].start, tt.start)
		}
		if ranges[i].end != tt.end {
			t.Errorf("end = %d; want %d", ranges[i].end, tt.end)
		}
	}
}
