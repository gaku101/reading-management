// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gaku101/my-portfolio/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/gaku101/my-portfolio/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockStore) AddAccountBalance(arg0 context.Context, arg1 db.AddAccountBalanceParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockStoreMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockStore)(nil).AddAccountBalance), arg0, arg1)
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateEntry mocks base method.
func (m *MockStore) CreateEntry(arg0 context.Context, arg1 db.CreateEntryParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockStoreMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockStore)(nil).CreateEntry), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockStore) CreatePost(arg0 context.Context, arg1 db.CreatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockStoreMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockStore)(nil).CreatePost), arg0, arg1)
}

// CreatePostCategory mocks base method.
func (m *MockStore) CreatePostCategory(arg0 context.Context, arg1 db.CreatePostCategoryParams) (db.PostCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostCategory", arg0, arg1)
	ret0, _ := ret[0].(db.PostCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePostCategory indicates an expected call of CreatePostCategory.
func (mr *MockStoreMockRecorder) CreatePostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostCategory", reflect.TypeOf((*MockStore)(nil).CreatePostCategory), arg0, arg1)
}

// CreatePostFavorite mocks base method.
func (m *MockStore) CreatePostFavorite(arg0 context.Context, arg1 db.CreatePostFavoriteParams) (db.PostFavorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.PostFavorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePostFavorite indicates an expected call of CreatePostFavorite.
func (mr *MockStoreMockRecorder) CreatePostFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostFavorite", reflect.TypeOf((*MockStore)(nil).CreatePostFavorite), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(arg0 context.Context, arg1 db.CreateTransferParams) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetAccountByOwner mocks base method.
func (m *MockStore) GetAccountByOwner(arg0 context.Context, arg1 string) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByOwner", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByOwner indicates an expected call of GetAccountByOwner.
func (mr *MockStoreMockRecorder) GetAccountByOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByOwner", reflect.TypeOf((*MockStore)(nil).GetAccountByOwner), arg0, arg1)
}

// GetAccountForUpdate mocks base method.
func (m *MockStore) GetAccountForUpdate(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockStoreMockRecorder) GetAccountForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountForUpdate), arg0, arg1)
}

// GetCategory mocks base method.
func (m *MockStore) GetCategory(arg0 context.Context, arg1 int64) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockStoreMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockStore)(nil).GetCategory), arg0, arg1)
}

// GetEntry mocks base method.
func (m *MockStore) GetEntry(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockStoreMockRecorder) GetEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockStore)(nil).GetEntry), arg0, arg1)
}

// GetPost mocks base method.
func (m *MockStore) GetPost(arg0 context.Context, arg1 int64) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockStoreMockRecorder) GetPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockStore)(nil).GetPost), arg0, arg1)
}

// GetPostCategory mocks base method.
func (m *MockStore) GetPostCategory(arg0 context.Context, arg1 int64) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostCategory indicates an expected call of GetPostCategory.
func (mr *MockStoreMockRecorder) GetPostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostCategory", reflect.TypeOf((*MockStore)(nil).GetPostCategory), arg0, arg1)
}

// GetPostFavorite mocks base method.
func (m *MockStore) GetPostFavorite(arg0 context.Context, arg1 db.GetPostFavoriteParams) (db.PostFavorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.PostFavorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostFavorite indicates an expected call of GetPostFavorite.
func (mr *MockStoreMockRecorder) GetPostFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostFavorite", reflect.TypeOf((*MockStore)(nil).GetPostFavorite), arg0, arg1)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(arg0 context.Context, arg1 int64) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// GetUserImage mocks base method.
func (m *MockStore) GetUserImage(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserImage", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserImage indicates an expected call of GetUserImage.
func (mr *MockStoreMockRecorder) GetUserImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserImage", reflect.TypeOf((*MockStore)(nil).GetUserImage), arg0, arg1)
}

// ListCategories mocks base method.
func (m *MockStore) ListCategories(arg0 context.Context) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", arg0)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCategories indicates an expected call of ListCategories.
func (mr *MockStoreMockRecorder) ListCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockStore)(nil).ListCategories), arg0)
}

// ListEntries mocks base method.
func (m *MockStore) ListEntries(arg0 context.Context, arg1 db.ListEntriesParams) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", arg0, arg1)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockStoreMockRecorder) ListEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockStore)(nil).ListEntries), arg0, arg1)
}

// ListFavoritePosts mocks base method.
func (m *MockStore) ListFavoritePosts(arg0 context.Context, arg1 db.ListFavoritePostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFavoritePosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFavoritePosts indicates an expected call of ListFavoritePosts.
func (mr *MockStoreMockRecorder) ListFavoritePosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFavoritePosts", reflect.TypeOf((*MockStore)(nil).ListFavoritePosts), arg0, arg1)
}

// ListMyPosts mocks base method.
func (m *MockStore) ListMyPosts(arg0 context.Context, arg1 db.ListMyPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMyPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMyPosts indicates an expected call of ListMyPosts.
func (mr *MockStoreMockRecorder) ListMyPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMyPosts", reflect.TypeOf((*MockStore)(nil).ListMyPosts), arg0, arg1)
}

// ListPosts mocks base method.
func (m *MockStore) ListPosts(arg0 context.Context, arg1 db.ListPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts.
func (mr *MockStoreMockRecorder) ListPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*MockStore)(nil).ListPosts), arg0, arg1)
}

// ListTransfers mocks base method.
func (m *MockStore) ListTransfers(arg0 context.Context, arg1 db.ListTransfersParams) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockStoreMockRecorder) ListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockStore)(nil).ListTransfers), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(arg0 context.Context, arg1 db.TransferTxParams) (db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockStore) UpdatePost(arg0 context.Context, arg1 db.UpdatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockStoreMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockStore)(nil).UpdatePost), arg0, arg1)
}

// UpdatePostCategory mocks base method.
func (m *MockStore) UpdatePostCategory(arg0 context.Context, arg1 db.UpdatePostCategoryParams) (db.PostCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePostCategory", arg0, arg1)
	ret0, _ := ret[0].(db.PostCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePostCategory indicates an expected call of UpdatePostCategory.
func (mr *MockStoreMockRecorder) UpdatePostCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePostCategory", reflect.TypeOf((*MockStore)(nil).UpdatePostCategory), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
