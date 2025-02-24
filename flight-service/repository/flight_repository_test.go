package repository

import (
	"errors"
	"flight-service/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// MockDB is a mock type for gorm.DB
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(value, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Model(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	callArgs := m.Called(append([]interface{}{query}, args...)...)
	return callArgs.Get(0).(*gorm.DB)
}

func (m *MockDB) Updates(values interface{}) *gorm.DB {
	args := m.Called(values)
	return args.Get(0).(*gorm.DB)
}

// FlightRepositoryTestSuite is a test suite for flightRepository
type FlightRepositoryTestSuite struct {
	suite.Suite
	mockDB  *MockDB
	logger  *zap.Logger
	repo    *flightRepository
	flight  *models.Flight
	flights []*models.Flight
}

// SetupTest sets up the test suite
func (suite *FlightRepositoryTestSuite) SetupTest() {
	suite.mockDB = new(MockDB)
	suite.logger = zap.NewNop()
	suite.repo = &flightRepository{db: &gorm.DB{}, Logger: suite.logger}
	suite.flight = &models.Flight{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FlightNumber: "123", Airline: "Airline", Origin: "Origin", Destination: "Destination", DepartureTime: "DepartureTime", ArrivalTime: "ArrivalTime"}
	suite.flights = []*models.Flight{suite.flight}
}

// TestCreate tests the Create method
func (suite *FlightRepositoryTestSuite) TestCreate() {
	suite.mockDB.On("Create", suite.flight).Return(&gorm.DB{Error: nil})
	flight, err := suite.repo.Create(suite.flight)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.flight, flight)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestCreateError tests the Create method with an error
func (suite *FlightRepositoryTestSuite) TestCreateError() {
	suite.mockDB.On("Create", suite.flight).Return(&gorm.DB{Error: errors.New("create error")})

	flight, err := suite.repo.Create(suite.flight)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), flight)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestDelete tests the Delete method
func (suite *FlightRepositoryTestSuite) TestDelete() {
	suite.mockDB.On("Delete", &models.Flight{}, uint(1)).Return(&gorm.DB{Error: nil})

	err := suite.repo.Delete(1)
	assert.NoError(suite.T(), err)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestDeleteError tests the Delete method with an error
func (suite *FlightRepositoryTestSuite) TestDeleteError() {
	suite.mockDB.On("Delete", &models.Flight{}, uint(1)).Return(&gorm.DB{Error: errors.New("delete error")})

	err := suite.repo.Delete(1)
	assert.Error(suite.T(), err)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestGet tests the Get method
func (suite *FlightRepositoryTestSuite) TestGet() {
	suite.mockDB.On("First", &models.Flight{}, uint(1)).Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Flight)
		*arg = *suite.flight
	})

	flight, err := suite.repo.Get(1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.flight, flight)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestGetError tests the Get method with an error
func (suite *FlightRepositoryTestSuite) TestGetError() {
	suite.mockDB.On("First", &models.Flight{}, uint(1)).Return(&gorm.DB{Error: errors.New("get error")})

	flight, err := suite.repo.Get(1)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), flight)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestGetAll tests the GetAll method
func (suite *FlightRepositoryTestSuite) TestGetAll() {
	suite.mockDB.On("Find", &suite.flights).Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]*models.Flight)
		*arg = suite.flights
	})

	flights, err := suite.repo.GetAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.flights, flights)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestGetAllError tests the GetAll method with an error
func (suite *FlightRepositoryTestSuite) TestGetAllError() {
	suite.mockDB.On("Find", &suite.flights).Return(&gorm.DB{Error: errors.New("get all error")})

	flights, err := suite.repo.GetAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), flights)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestUpdate tests the Update method
func (suite *FlightRepositoryTestSuite) TestUpdate() {
	suite.mockDB.On("Model", &models.Flight{}).Return(&gorm.DB{Error: nil}).Once()
	suite.mockDB.On("Where", "id = ?", uint(1)).Return(&gorm.DB{Error: nil}).Once()
	suite.mockDB.On("Updates", suite.flight).Return(&gorm.DB{Error: nil}).Once()

	flight, err := suite.repo.Update(1, suite.flight)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.flight, flight)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestUpdateError tests the Update method with an error
func (suite *FlightRepositoryTestSuite) TestUpdateError() {
	suite.mockDB.On("Model", &models.Flight{}).Return(&gorm.DB{Error: nil}).Once()
	suite.mockDB.On("Where", "id = ?", uint(1)).Return(&gorm.DB{Error: nil}).Once()
	suite.mockDB.On("Updates", suite.flight).Return(&gorm.DB{Error: errors.New("update error")}).Once()

	flight, err := suite.repo.Update(1, suite.flight)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), flight)
	suite.mockDB.AssertExpectations(suite.T())
}

// TestFlightRepositoryTestSuite runs the test suite
func TestFlightRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(FlightRepositoryTestSuite))
}
