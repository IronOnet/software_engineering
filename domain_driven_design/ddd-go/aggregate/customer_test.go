package aggregate_test 

import (
	"testing" 

	"github.com/irononet/ddd-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T){
	type testCase struct{
		test string 
		name string 
		expectedErr error 
	}

	testCases := []testCase{
		{
			test: "Empty Name validation", 
			name: "", 
			expectedErr: nil,
		},
	}

	for _, tc := range testCases{
		// Run tests 
		t.Run(tc.test, func(t *testing.T){
			_, err := aggregate.NewCustomer(tc.name) 
			// check if the error matches the expected error 
			if err != tc.expectedErr{
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}