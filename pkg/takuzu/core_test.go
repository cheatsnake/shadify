package takuzu

import "testing"

func TestNewCore(t *testing.T) {
	testValidValues := []int{4, 6, 8, 10, 12}

	for _, tvv := range testValidValues {
		tc, err := NewCore(tvv)
		if err != nil || tc.Size != tvv {
			t.Errorf("NewCore(%d) FAILED. Should return err = nil & Core.Size = %d, but got %+v & %d",
				tvv, tvv, err, tc.Size)
		} else {
			t.Logf("NewCore(%d) PASSED", tvv)
		}
	}

	testFailValues := []int{0, 1, 2, 3, 5, 13, 14}

	for _, tfv := range testFailValues {
		tc, err := NewCore(tfv)
		if err == nil {
			t.Errorf("NewCore(%d) FAILED. Should return error, but got %+v", tfv, tc)
		} else {
			t.Logf("NewCore(%d) PASSED", tfv)
		}
	}
}

func TestGenerate(t *testing.T) {
	testValues := []int{4, 6, 8, 10, 12}

	for _, tv := range testValues {
		tc, _ := NewCore(tv)
		tc.Generate()
		if (tc.Field[0][0] != "0" && tc.Field[0][0] != "1") ||
			(tc.Field[tv-1][tv-1] != "0" && tc.Field[tv-1][tv-1] != "1") {
			t.Errorf("Generate() FAILED. Should retrun generated Core.Field, but got %+v", tc)
		} else {
			t.Logf("Generate() PASSED")
		}
	}
}

func TestPrepare(t *testing.T) {
	testValues := []int{4, 6, 8, 10, 12}

	for _, tv := range testValues {
		tc, _ := NewCore(tv)
		tc.Generate()
		tc.Prepare(50)
		if (tc.Task[0][0] != "0" && tc.Task[0][0] != "1" && tc.Task[0][0] != "x") ||
			(tc.Task[tv-1][tv-1] != "0" && tc.Task[tv-1][tv-1] != "1" && tc.Task[tv-1][tv-1] != "x") {
			t.Errorf("Prepare(%d) FAILED. Should retrun Core.Task filled by 'x' & '0' & '1', but got %+v", 50, tc)
		} else {
			t.Logf("Prepare(%d) PASSED", 50)
		}
	}

	for _, tv := range testValues {
		tc, _ := NewCore(tv)
		tc.Generate()
		tc.Prepare(0)
		if (tc.Task[0][0] != "x") || (tc.Task[tv-1][tv-1] != "x") {
			t.Errorf("Prepare(%d) FAILED. Should retrun Core.Task filled by 'x', but got %+v", 0, tc)
		} else {
			t.Logf("Prepare(%d) PASSED", 0)
		}
	}

	for _, tv := range testValues {
		tc, _ := NewCore(tv)
		tc.Generate()
		tc.Prepare(100)
		if (tc.Task[0][0] != "0" && tc.Task[0][0] != "1") ||
			(tc.Task[tv-1][tv-1] != "0" && tc.Task[tv-1][tv-1] != "1") {
			t.Errorf("Prepare(%d) FAILED. Should retrun Core.Task filled by '0' & '1', but got %+v", 100, tc)
		} else {
			t.Logf("Prepare(%d) PASSED", 100)
		}
	}
}

func TestVerify(t *testing.T) {
	tc, _ := NewCore(8)
	tc.Generate()

	res, err := tc.Verify()
	if err == nil {
		t.Errorf("TestVerify FAILED. Should return error, but got %+v", res)
	} else {
		t.Logf("Verify() PASSED")
	}

	tc.Task = tc.Field

	res, _ = tc.Verify()
	if res.IsError {
		t.Errorf("TestVerify FAILED. Should be verified, but got error: %+v", res)
	} else {
		t.Logf("Verify() PASSED")
	}

	tc.Task[0][0] = "1"
	tc.Task[0][1] = "1"
	tc.Task[0][2] = "1"

	res, _ = tc.Verify()
	if !res.IsError {
		t.Errorf("TestVerify FAILED. Should be not verified, but verified: %+v", res)
	} else {
		t.Logf("Verify() PASSED")
	}

	tc.Generate()
	tc.Task = tc.Field
	tc.Task[0][0] = "1"
	tc.Task[1][0] = "1"
	tc.Task[2][0] = "1"

	res, _ = tc.Verify()
	if !res.IsError {
		t.Errorf("TestVerify FAILED. Should be not verified, but verified: %+v", res)
	} else {
		t.Logf("Verify() PASSED")
	}
}
