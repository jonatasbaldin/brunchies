package function

import "testing"

func Test_Week_Equal(t *testing.T) {
	week1 := Week{
		Year: 2019,
		Week: 1,
	}

	week2 := Week{
		Year: 2019,
		Week: 2,
	}

	if week1.Equal(week2) {
		t.Errorf("expected %v, got %v", false, true)
	}
}

func Test_Week_GetSpecificWeek(t *testing.T) {
	weeks := []Week{
		{
			Year: 2019,
			Week: 1,
		},
	}

	week := getSpecificWeek(weeks, Week{Year: 2019, Week: 1})
	if week != 0 {
		t.Errorf("expected %d, got %d", 0, week)
	}
}
