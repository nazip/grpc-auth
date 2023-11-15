package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/nazip/grpc-auth/internal/repository.UserRepository -o ./mocks/user_repository_minimock.go -n UserRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	model "github.com/nazip/grpc-auth/internal/models/service"
)

// UserRepositoryMock implements repository.UserRepository
type UserRepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, req *model.User) (u1 uint64, err error)
	inspectFuncCreate   func(ctx context.Context, req *model.User)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mUserRepositoryMockCreate

	funcDelete          func(ctx context.Context, id uint64) (err error)
	inspectFuncDelete   func(ctx context.Context, id uint64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mUserRepositoryMockDelete

	funcGet          func(ctx context.Context, id uint64) (up1 *model.User, err error)
	inspectFuncGet   func(ctx context.Context, id uint64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUserRepositoryMockGet

	funcUpdate          func(ctx context.Context, req *model.User) (err error)
	inspectFuncUpdate   func(ctx context.Context, req *model.User)
	afterUpdateCounter  uint64
	beforeUpdateCounter uint64
	UpdateMock          mUserRepositoryMockUpdate
}

// NewUserRepositoryMock returns a mock for repository.UserRepository
func NewUserRepositoryMock(t minimock.Tester) *UserRepositoryMock {
	m := &UserRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mUserRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*UserRepositoryMockCreateParams{}

	m.DeleteMock = mUserRepositoryMockDelete{mock: m}
	m.DeleteMock.callArgs = []*UserRepositoryMockDeleteParams{}

	m.GetMock = mUserRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*UserRepositoryMockGetParams{}

	m.UpdateMock = mUserRepositoryMockUpdate{mock: m}
	m.UpdateMock.callArgs = []*UserRepositoryMockUpdateParams{}

	return m
}

type mUserRepositoryMockCreate struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockCreateExpectation
	expectations       []*UserRepositoryMockCreateExpectation

	callArgs []*UserRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// UserRepositoryMockCreateExpectation specifies expectation struct of the UserRepository.Create
type UserRepositoryMockCreateExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockCreateParams
	results *UserRepositoryMockCreateResults
	Counter uint64
}

// UserRepositoryMockCreateParams contains parameters of the UserRepository.Create
type UserRepositoryMockCreateParams struct {
	ctx context.Context
	req *model.User
}

// UserRepositoryMockCreateResults contains results of the UserRepository.Create
type UserRepositoryMockCreateResults struct {
	u1  uint64
	err error
}

// Expect sets up expected params for UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Expect(ctx context.Context, req *model.User) *mUserRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &UserRepositoryMockCreateParams{ctx, req}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Inspect(f func(ctx context.Context, req *model.User)) *mUserRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Return(u1 uint64, err error) *UserRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &UserRepositoryMockCreateResults{u1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the UserRepository.Create method
func (mmCreate *mUserRepositoryMockCreate) Set(f func(ctx context.Context, req *model.User) (u1 uint64, err error)) *UserRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the UserRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the UserRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the UserRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mUserRepositoryMockCreate) When(ctx context.Context, req *model.User) *UserRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	expectation := &UserRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &UserRepositoryMockCreateParams{ctx, req},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Create return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockCreateExpectation) Then(u1 uint64, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockCreateResults{u1, err}
	return e.mock
}

// Create implements repository.UserRepository
func (mmCreate *UserRepositoryMock) Create(ctx context.Context, req *model.User) (u1 uint64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, req)
	}

	mm_params := &UserRepositoryMockCreateParams{ctx, req}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := UserRepositoryMockCreateParams{ctx, req}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("UserRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the UserRepositoryMock.Create")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, req)
	}
	mmCreate.t.Fatalf("Unexpected call to UserRepositoryMock.Create. %v %v", ctx, req)
	return
}

// CreateAfterCounter returns a count of finished UserRepositoryMock.Create invocations
func (mmCreate *UserRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of UserRepositoryMock.Create invocations
func (mmCreate *UserRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mUserRepositoryMockCreate) Calls() []*UserRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*UserRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.Create")
	}
}

type mUserRepositoryMockDelete struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockDeleteExpectation
	expectations       []*UserRepositoryMockDeleteExpectation

	callArgs []*UserRepositoryMockDeleteParams
	mutex    sync.RWMutex
}

// UserRepositoryMockDeleteExpectation specifies expectation struct of the UserRepository.Delete
type UserRepositoryMockDeleteExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockDeleteParams
	results *UserRepositoryMockDeleteResults
	Counter uint64
}

// UserRepositoryMockDeleteParams contains parameters of the UserRepository.Delete
type UserRepositoryMockDeleteParams struct {
	ctx context.Context
	id  uint64
}

// UserRepositoryMockDeleteResults contains results of the UserRepository.Delete
type UserRepositoryMockDeleteResults struct {
	err error
}

// Expect sets up expected params for UserRepository.Delete
func (mmDelete *mUserRepositoryMockDelete) Expect(ctx context.Context, id uint64) *mUserRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UserRepositoryMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &UserRepositoryMockDeleteParams{ctx, id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Delete
func (mmDelete *mUserRepositoryMockDelete) Inspect(f func(ctx context.Context, id uint64)) *mUserRepositoryMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by UserRepository.Delete
func (mmDelete *mUserRepositoryMockDelete) Return(err error) *UserRepositoryMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UserRepositoryMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &UserRepositoryMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the UserRepository.Delete method
func (mmDelete *mUserRepositoryMockDelete) Set(f func(ctx context.Context, id uint64) (err error)) *UserRepositoryMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the UserRepository.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the UserRepository.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the UserRepository.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mUserRepositoryMockDelete) When(ctx context.Context, id uint64) *UserRepositoryMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserRepositoryMock.Delete mock is already set by Set")
	}

	expectation := &UserRepositoryMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &UserRepositoryMockDeleteParams{ctx, id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Delete return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockDeleteExpectation) Then(err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockDeleteResults{err}
	return e.mock
}

// Delete implements repository.UserRepository
func (mmDelete *UserRepositoryMock) Delete(ctx context.Context, id uint64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, id)
	}

	mm_params := &UserRepositoryMockDeleteParams{ctx, id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := UserRepositoryMockDeleteParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("UserRepositoryMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the UserRepositoryMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, id)
	}
	mmDelete.t.Fatalf("Unexpected call to UserRepositoryMock.Delete. %v %v", ctx, id)
	return
}

// DeleteAfterCounter returns a count of finished UserRepositoryMock.Delete invocations
func (mmDelete *UserRepositoryMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of UserRepositoryMock.Delete invocations
func (mmDelete *UserRepositoryMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mUserRepositoryMockDelete) Calls() []*UserRepositoryMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*UserRepositoryMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.Delete")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.Delete")
	}
}

type mUserRepositoryMockGet struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockGetExpectation
	expectations       []*UserRepositoryMockGetExpectation

	callArgs []*UserRepositoryMockGetParams
	mutex    sync.RWMutex
}

// UserRepositoryMockGetExpectation specifies expectation struct of the UserRepository.Get
type UserRepositoryMockGetExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockGetParams
	results *UserRepositoryMockGetResults
	Counter uint64
}

// UserRepositoryMockGetParams contains parameters of the UserRepository.Get
type UserRepositoryMockGetParams struct {
	ctx context.Context
	id  uint64
}

// UserRepositoryMockGetResults contains results of the UserRepository.Get
type UserRepositoryMockGetResults struct {
	up1 *model.User
	err error
}

// Expect sets up expected params for UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Expect(ctx context.Context, id uint64) *mUserRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserRepositoryMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &UserRepositoryMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Inspect(f func(ctx context.Context, id uint64)) *mUserRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by UserRepository.Get
func (mmGet *mUserRepositoryMockGet) Return(up1 *model.User, err error) *UserRepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserRepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UserRepositoryMockGetResults{up1, err}
	return mmGet.mock
}

// Set uses given function f to mock the UserRepository.Get method
func (mmGet *mUserRepositoryMockGet) Set(f func(ctx context.Context, id uint64) (up1 *model.User, err error)) *UserRepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the UserRepository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the UserRepository.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the UserRepository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUserRepositoryMockGet) When(ctx context.Context, id uint64) *UserRepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserRepositoryMock.Get mock is already set by Set")
	}

	expectation := &UserRepositoryMockGetExpectation{
		mock:   mmGet.mock,
		params: &UserRepositoryMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Get return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockGetExpectation) Then(up1 *model.User, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockGetResults{up1, err}
	return e.mock
}

// Get implements repository.UserRepository
func (mmGet *UserRepositoryMock) Get(ctx context.Context, id uint64) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := &UserRepositoryMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := UserRepositoryMockGetParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("UserRepositoryMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the UserRepositoryMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to UserRepositoryMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished UserRepositoryMock.Get invocations
func (mmGet *UserRepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UserRepositoryMock.Get invocations
func (mmGet *UserRepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUserRepositoryMockGet) Calls() []*UserRepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UserRepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.Get")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.Get")
	}
}

type mUserRepositoryMockUpdate struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockUpdateExpectation
	expectations       []*UserRepositoryMockUpdateExpectation

	callArgs []*UserRepositoryMockUpdateParams
	mutex    sync.RWMutex
}

// UserRepositoryMockUpdateExpectation specifies expectation struct of the UserRepository.Update
type UserRepositoryMockUpdateExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockUpdateParams
	results *UserRepositoryMockUpdateResults
	Counter uint64
}

// UserRepositoryMockUpdateParams contains parameters of the UserRepository.Update
type UserRepositoryMockUpdateParams struct {
	ctx context.Context
	req *model.User
}

// UserRepositoryMockUpdateResults contains results of the UserRepository.Update
type UserRepositoryMockUpdateResults struct {
	err error
}

// Expect sets up expected params for UserRepository.Update
func (mmUpdate *mUserRepositoryMockUpdate) Expect(ctx context.Context, req *model.User) *mUserRepositoryMockUpdate {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("UserRepositoryMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &UserRepositoryMockUpdateExpectation{}
	}

	mmUpdate.defaultExpectation.params = &UserRepositoryMockUpdateParams{ctx, req}
	for _, e := range mmUpdate.expectations {
		if minimock.Equal(e.params, mmUpdate.defaultExpectation.params) {
			mmUpdate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdate.defaultExpectation.params)
		}
	}

	return mmUpdate
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Update
func (mmUpdate *mUserRepositoryMockUpdate) Inspect(f func(ctx context.Context, req *model.User)) *mUserRepositoryMockUpdate {
	if mmUpdate.mock.inspectFuncUpdate != nil {
		mmUpdate.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Update")
	}

	mmUpdate.mock.inspectFuncUpdate = f

	return mmUpdate
}

// Return sets up results that will be returned by UserRepository.Update
func (mmUpdate *mUserRepositoryMockUpdate) Return(err error) *UserRepositoryMock {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("UserRepositoryMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &UserRepositoryMockUpdateExpectation{mock: mmUpdate.mock}
	}
	mmUpdate.defaultExpectation.results = &UserRepositoryMockUpdateResults{err}
	return mmUpdate.mock
}

// Set uses given function f to mock the UserRepository.Update method
func (mmUpdate *mUserRepositoryMockUpdate) Set(f func(ctx context.Context, req *model.User) (err error)) *UserRepositoryMock {
	if mmUpdate.defaultExpectation != nil {
		mmUpdate.mock.t.Fatalf("Default expectation is already set for the UserRepository.Update method")
	}

	if len(mmUpdate.expectations) > 0 {
		mmUpdate.mock.t.Fatalf("Some expectations are already set for the UserRepository.Update method")
	}

	mmUpdate.mock.funcUpdate = f
	return mmUpdate.mock
}

// When sets expectation for the UserRepository.Update which will trigger the result defined by the following
// Then helper
func (mmUpdate *mUserRepositoryMockUpdate) When(ctx context.Context, req *model.User) *UserRepositoryMockUpdateExpectation {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("UserRepositoryMock.Update mock is already set by Set")
	}

	expectation := &UserRepositoryMockUpdateExpectation{
		mock:   mmUpdate.mock,
		params: &UserRepositoryMockUpdateParams{ctx, req},
	}
	mmUpdate.expectations = append(mmUpdate.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Update return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockUpdateExpectation) Then(err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockUpdateResults{err}
	return e.mock
}

// Update implements repository.UserRepository
func (mmUpdate *UserRepositoryMock) Update(ctx context.Context, req *model.User) (err error) {
	mm_atomic.AddUint64(&mmUpdate.beforeUpdateCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdate.afterUpdateCounter, 1)

	if mmUpdate.inspectFuncUpdate != nil {
		mmUpdate.inspectFuncUpdate(ctx, req)
	}

	mm_params := &UserRepositoryMockUpdateParams{ctx, req}

	// Record call args
	mmUpdate.UpdateMock.mutex.Lock()
	mmUpdate.UpdateMock.callArgs = append(mmUpdate.UpdateMock.callArgs, mm_params)
	mmUpdate.UpdateMock.mutex.Unlock()

	for _, e := range mmUpdate.UpdateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdate.UpdateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdate.UpdateMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdate.UpdateMock.defaultExpectation.params
		mm_got := UserRepositoryMockUpdateParams{ctx, req}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdate.t.Errorf("UserRepositoryMock.Update got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdate.UpdateMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdate.t.Fatal("No results are set for the UserRepositoryMock.Update")
		}
		return (*mm_results).err
	}
	if mmUpdate.funcUpdate != nil {
		return mmUpdate.funcUpdate(ctx, req)
	}
	mmUpdate.t.Fatalf("Unexpected call to UserRepositoryMock.Update. %v %v", ctx, req)
	return
}

// UpdateAfterCounter returns a count of finished UserRepositoryMock.Update invocations
func (mmUpdate *UserRepositoryMock) UpdateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.afterUpdateCounter)
}

// UpdateBeforeCounter returns a count of UserRepositoryMock.Update invocations
func (mmUpdate *UserRepositoryMock) UpdateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.beforeUpdateCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Update.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdate *mUserRepositoryMockUpdate) Calls() []*UserRepositoryMockUpdateParams {
	mmUpdate.mutex.RLock()

	argCopy := make([]*UserRepositoryMockUpdateParams, len(mmUpdate.callArgs))
	copy(argCopy, mmUpdate.callArgs)

	mmUpdate.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateDone returns true if the count of the Update invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockUpdateDone() bool {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockUpdateInspect() {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Update with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		if m.UpdateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.Update")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Update with params: %#v", *m.UpdateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.Update")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteInspect()

		m.MinimockGetInspect()

		m.MinimockUpdateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone() &&
		m.MinimockUpdateDone()
}
