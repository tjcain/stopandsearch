package postgres

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	want := "force2018-01"
	got := generateID("Force", "2018-01")
	if got != want {
		t.Errorf("generateID got %s want %s", got, want)
	}
}

func TestNormalizeEthnicity(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Black",
			input: "Black/African/Caribbean/Black British - Caribbean",
			want:  ethnicityBlack,
		},
		{
			name:  "White",
			input: "White - Any other White background",
			want:  ethnicityWhite,
		},
		{
			name:  "Asian",
			input: "Asian/Asian British - Pakistani",
			want:  ethnicityAsian,
		},
		{
			name:  "Other",
			input: "Other ethnic group - Not stated",
			want:  ethnicityOther,
		},
		{
			name:  "Mixed",
			input: "Mixed/Multiple ethnic groups - White and Asian",
			want:  ethnicityMixed,
		},
		{
			name:  "Not Stated",
			input: "",
			want:  ethnicityNotStated,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := normalizeEthnicity(tc.input)
			if got != tc.want {
				t.Errorf("normalizeEthnicity %s got %s want %s", tc.name, got, tc.want)
			}
		})
	}
}
