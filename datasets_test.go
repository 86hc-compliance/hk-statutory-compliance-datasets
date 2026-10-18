package datasets

import "testing"

func TestDatasets(t *testing.T) {
	banks := GetLicensedBanks()
	if len(banks) == 0 || banks[0].ClearingCode != "004" {
		t.Fatalf("expected HSBC clearing code 004")
	}
	if CalculateNAR1Fee(30) != 105 || CalculateNAR1Fee(300) != 3480 {
		t.Fatalf("unexpected NAR1 fee calculation")
	}
}
