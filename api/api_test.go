package api

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeDB struct {
	received    string
	returned    string
	returnedErr error
}

func (f *fakeDB) GetThing(r string) (string, error) {
	f.received = r
	return f.returned, f.returnedErr
}

func devClient(t *testing.T) {

}

// don't test api client.  test the actual endpoint
func TestGetThing(t *testing.T) {
	type testCase struct {
		name               string
		requestedID        string
		dbReturned         string
		dbReturnedErr      error
		expectedBody       string
		expectedStatusCode int
	}
	tcs := []testCase{
		{
			name:               "SuccessfulRequest",
			requestedID:        "xyz",
			dbReturned:         "thing1",
			dbReturnedErr:      nil,
			expectedBody:       "thing1_processed",
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "SimulateDBError",
			requestedID:        "xyz",
			dbReturned:         "",
			dbReturnedErr:      fmt.Errorf("Error in DB"),
			expectedBody:       "Error in DB",
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name, func(t *testing.T) {
			fDb := fakeDB{
				returned:    tc.dbReturned,
				returnedErr: tc.dbReturnedErr,
			}
			s := httptest.NewServer(makeGetThing(&fDb))
			resp, err := http.Get(fmt.Sprintf("%s/%s", s.URL, tc.requestedID))
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatusCode, resp.StatusCode)

			// Show that the parameter was correctly parsed and sent to DB
			assert.Equal(t, tc.requestedID, fDb.received)

			// Show that api responds as expected
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedBody, string(body))

		})
	}

}
