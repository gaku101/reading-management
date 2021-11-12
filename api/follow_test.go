package api

import (
	"bytes"
	"database/sql"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/gaku101/my-portfolio/db/mock"
	db "github.com/gaku101/my-portfolio/db/sqlc"
	"github.com/gaku101/my-portfolio/token"
	"github.com/gaku101/my-portfolio/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

func TestCreateFollowAPI(t *testing.T) {
	user, _ := randomUser(t)
	follow := randomFollow()

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"followingId": follow.FollowingID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)

				arg := db.CreateFollowParams{
					FollowingID: follow.FollowingID,
					FollowerID:  user.ID,
				}
				store.EXPECT().
					CreateFollow(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(follow, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchFollow(t, recorder.Body, follow)
			},
		},
		{
			name: "NoAuthorization",
			body: gin.H{
				"followingId": follow.FollowingID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateFollow(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"followingId": follow.FollowingID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)
				store.EXPECT().
					CreateFollow(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Follow{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"followingId": user.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)

				arg := db.CreateFollowParams{
					FollowingID: user.ID,
					FollowerID:  user.ID,
				}
				store.EXPECT().
					CreateFollow(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/follow"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
func randomFollow() db.Follow {
	return db.Follow{
		ID:          util.RandomInt(1, 1000),
		FollowingID: util.RandomInt(1, 1000),
		FollowerID:  util.RandomInt(1, 1000),
	}
}
func requireBodyMatchFollow(t *testing.T, body *bytes.Buffer, follow db.Follow) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotFollow db.Follow
	err = json.Unmarshal(data, &gotFollow)
	require.NoError(t, err)
	require.Equal(t, follow, gotFollow)
}
