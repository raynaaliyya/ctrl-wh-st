package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyAdmin = []entity.WarehouseTeam{
	{
		ID:       "1",
		Name:     "Rayna",
		Email:    "rayna@mail.com",
		Password: "12345",
		Phone:    "08123456789",
		Photo:    "photo.jpg",
	},
	{
		ID:       "2",
		Name:     "Siti",
		Email:    "siti@mail.com",
		Password: "22345",
		Phone:    "08223456789",
		Photo:    "photo.jpg",
	},
	{
		ID:       "3",
		Name:     "Aliyya",
		Email:    "aliyya@mail.com",
		Password: "32345",
		Phone:    "08323456789",
		Photo:    "photo.jpg",
	},
}

var dummyAdminRes = []entity.EmployeeResponse{
	{
		ID:    "1",
		Name:  "Rayna",
		Email: "rayna@mail.com",
		Phone: "08123456789",
		Photo: "photo.jpg",
	},
	{
		ID:    "2",
		Name:  "Siti",
		Email: "siti@mail.com",
		Phone: "08223456789",
		Photo: "photo.jpg",
	},
	{
		ID:    "3",
		Name:  "Aliyya",
		Email: "aliyya@mail.com",
		Phone: "08323456789",
		Photo: "photo.jpg",
	},
}

type repoMock struct {
	mock.Mock
}

type WarehouseTeamUsecaseTestSuite struct {
	repoMock *repoMock
	suite.Suite
}

func (r *repoMock) GetAll() any {
	arg := r.Called()
	if arg.Get(0) == nil {
		return []entity.EmployeeResponse{}
		// return "no data"
	}
	return arg.Get(0).([]entity.EmployeeResponse)
}

func (r *repoMock) GetById(id string) any {
	arg := r.Called(id)
	if arg.Get(0) == nil {
		return "employee not found"
	}
	return arg.Get(0)
}

func (r *repoMock) GetByEmail(email string) (*entity.WarehouseTeam, error) {
	return &entity.WarehouseTeam{}, nil
}

func (r *repoMock) Create(newEmployee *entity.WarehouseTeam) string {
	arg := r.Called(newEmployee)
	if arg[0] != nil {
		return arg.String(0)
	}
	return "admin created"
}

func (r *repoMock) Update(employee *entity.WarehouseTeam) string {
	arg := r.Called(employee)
	if arg[0] != nil {
		return arg.String(0)
	}
	return "employee updated"
}

func (r *repoMock) Delete(id string) string {
	arg := r.Called(id)
	if arg[0] != nil {
		return arg.String(0)
	}
	return "employee deleted"
}

func (suite *WarehouseTeamUsecaseTestSuite) TestRegisterWhTeam_Success() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("Create", &dummyAdmin[0]).Return("admin created")

	adminWh := whTeamUc.Register(&dummyAdmin[0])

	assert.Equal(suite.T(), "admin created", adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) TestRegisterWhTeam_Failed() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("Create", &dummyAdmin[0]).Return("failed to create employee")

	adminWh := whTeamUc.Register(&dummyAdmin[0])

	assert.Equal(suite.T(), "failed to create employee", adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) TestFindEmployees_Success() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("GetAll").Return(dummyAdminRes)

	adminWh := whTeamUc.FindEmployees()
	adminWhs := adminWh.([]entity.EmployeeResponse)

	assert.Equal(suite.T(), dummyAdminRes, adminWh)
	assert.Equal(suite.T(), len(dummyAdminRes), len(adminWhs))
}

func (suite *WarehouseTeamUsecaseTestSuite) TestFindEmployees_Failed() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("GetAll").Return([]entity.EmployeeResponse{})
	// suite.repoMock.On("GetAll").Return("no data")

	adminWh := whTeamUc.FindEmployees()
	adminWhs := adminWh.([]entity.EmployeeResponse)

	// assert.Equal(suite.T(),"no data", adminWh)
	assert.Equal(suite.T(), 0, len(adminWhs))
	assert.Empty(suite.T(), adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) TestFindEmployeeById_Success() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("GetById", "1").Return(dummyAdminRes[0].Photo)

	adminWh := whTeamUc.FindEmployeeById("1")

	assert.Equal(suite.T(), dummyAdminRes[0].Photo, adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) TestFindEmployeeById_Failed() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("GetById", "5").Return("no data")

	adminWh := whTeamUc.FindEmployeeById("5")

	assert.Equal(suite.T(), "no data", adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) TestEditEmployee_Success() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("Update", &dummyAdmin[0]).Return("admin updated")

	adminWh := whTeamUc.Edit(&dummyAdmin[0])

	assert.Equal(suite.T(), "admin updated", adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) TestEditEmployee_Failed() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("Update", &dummyAdmin[0]).Return("failed to update employee")

	adminWh := whTeamUc.Edit(&dummyAdmin[0])

	assert.Equal(suite.T(), "failed to update employee", adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) TestUnregEmployee_Success() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("Delete", "1").Return("admin deleted")

	adminWh := whTeamUc.Unreg("1")

	assert.Equal(suite.T(), "admin deleted", adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) TestUnregEmployee_Failed() {
	whTeamUc := NewWarehouseTeamUsecase(suite.repoMock)
	suite.repoMock.On("Delete", "5").Return("failed to delete employee")

	adminWh := whTeamUc.Unreg("5")

	assert.Equal(suite.T(), "failed to delete employee", adminWh)
}

func (suite *WarehouseTeamUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func TestWarehouseTeamUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(WarehouseTeamUsecaseTestSuite))
}
