package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go_inven_ctrl/entity"

	// "go_inven_ctrl/repository"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyAdmin = []entity.WarehouseTeam{
	{
		ID:       "1",
		Name:     "Rayna",
		Email:    "rayna@mail.com",
		Password: "12345",
		Phone:    "08123456789",
		Photo:    "Screenshot (598).png",
	},
	{
		ID:       "2",
		Name:     "Siti",
		Email:    "siti@mail.com",
		Password: "22345",
		Phone:    "08223456789",
		Photo:    "Screenshot (599).png",
	},
	{
		ID:       "3",
		Name:     "Aliyya",
		Email:    "aliyya@mail.com",
		Password: "32345",
		Phone:    "08323456789",
		Photo:    "Screenshot (600).png",
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

type WarehouseTeamRepoTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *WarehouseTeamRepoTestSuite) TestGetAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo"})

	for _, v := range dummyAdminRes {
		rows.AddRow(v.ID, v.Name, v.Email, v.Phone, v.Photo)
	}

	suite.mockSql.ExpectQuery("SELECT id, name, email, phone, photo FROM admin_wh").WillReturnRows(rows)

	repo := NewWarehouseTeamRepo(suite.mockDb)
	expected := dummyAdminRes
	actual := repo.GetAll().([]entity.EmployeeResponse)

	assert.Equal(suite.T(), expected, actual)
	assert.Equal(suite.T(), len(dummyAdmin), len(actual))
	assert.Equal(suite.T(), "2", actual[1].ID)
}

func (suite *WarehouseTeamRepoTestSuite) TestGetAllScan_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo"})
	rows.AddRow(nil, "Rayna", "rayna@mail.com", "08123456789", "photo.jpg")

	suite.mockSql.ExpectQuery("SELECT id, name, email, phone, photo FROM admin_wh").WillReturnRows(rows)

	repo := NewWarehouseTeamRepo(suite.mockDb)
	actual := repo.GetAll()

	expected := []entity.EmployeeResponse{
		{ID: "", Name: "", Email: "", Phone: "", Photo: ""},
	}
	assert.Equal(suite.T(), expected, actual)
}

func (suite *WarehouseTeamRepoTestSuite) TestGetAll_Empty() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "photo"})
	suite.mockSql.ExpectQuery("SELECT id, name, email, phone, photo FROM admin_wh").WillReturnRows(rows)

	repo := NewWarehouseTeamRepo(suite.mockDb)
	actual := repo.GetAll()

	assert.Equal(suite.T(), "no data", actual)
}

func (suite *WarehouseTeamRepoTestSuite) TestGetById_Success() {
	expectedPhoto := "Screenshot (600).png"

	suite.mockSql.ExpectQuery("SELECT photo FROM admin_wh WHERE id = ?").WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"photo"}).AddRow(expectedPhoto))

	repo := NewWarehouseTeamRepo(suite.mockDb)
	actualPhoto := repo.GetById("1")

	assert.Equal(suite.T(), expectedPhoto, actualPhoto)
}

func (suite *WarehouseTeamRepoTestSuite) TestGetById_Failed() {
	suite.mockSql.ExpectQuery("SELECT photo FROM admin_wh WHERE id = ?").WithArgs("employee not found").WillReturnError(errors.New("Failed to get photo"))

	repo := NewWarehouseTeamRepo(suite.mockDb)
	actualPhoto := repo.GetById("invalid-id")

	assert.NotNil(suite.T(), actualPhoto)
	assert.Equal(suite.T(), "employee not found", actualPhoto)
}

func (suite *WarehouseTeamRepoTestSuite) TestGetByEmail_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "password", "email", "phone", "photo"}).AddRow(dummyAdmin[0].ID, dummyAdmin[0].Name, dummyAdmin[0].Email, dummyAdmin[0].Password, dummyAdmin[0].Phone, dummyAdmin[0].Photo)

	suite.mockSql.ExpectQuery("SELECT id, name, email, password, phone, photo FROM admin_wh WHERE email = ?").WithArgs("rayna@mail.com").WillReturnRows(rows)

	repo := NewWarehouseTeamRepo(suite.mockDb)
	expected := &dummyAdmin[0]
	actual, err := repo.GetByEmail("rayna@mail.com")

	assert.Equal(suite.T(), expected, actual)
	assert.Nil(suite.T(), err)
}

func (suite *WarehouseTeamRepoTestSuite) TestGetByEmail_Failed() {
	suite.mockSql.ExpectQuery("SELECT id, name, email, password, phone, photo FROM admin_wh WHERE email = ?").WillReturnError(errors.New("employee not found"))

	repo := NewWarehouseTeamRepo(suite.mockDb)
	expectedError := errors.New("employee not found")
	actual, err := repo.GetByEmail("rayna@mail.com")

	assert.Nil(suite.T(), actual)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *WarehouseTeamRepoTestSuite) TestWarehouseTeamCreate_Success() {
	newWarehouseTeam := dummyAdmin[0]
	suite.mockSql.ExpectExec("INSERT INTO admin_wh\\(id, name, email, password, phone, photo\\) VALUES").WithArgs(
		newWarehouseTeam.ID,
		newWarehouseTeam.Name,
		newWarehouseTeam.Email,
		newWarehouseTeam.Password,
		newWarehouseTeam.Phone,
		newWarehouseTeam.Photo,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	warehouseTeamRepository := NewWarehouseTeamRepo(suite.mockDb)
	actual := warehouseTeamRepository.Create(&newWarehouseTeam)
	assert.Equal(suite.T(), "New employee created successfully", actual)
}

func (suite *WarehouseTeamRepoTestSuite) TestWarehouseTeamCreate_Failed() {
	newWarehouseTeam := dummyAdmin[0]
	suite.mockSql.ExpectExec("INSERT INTO admin_wh\\(id, name, email, password, phone, photo\\) VALUES").WillReturnError(errors.New("failed to create employee"))

	warehouseTeamRepository := NewWarehouseTeamRepo(suite.mockDb)
	actual := warehouseTeamRepository.Create(&newWarehouseTeam)
	assert.Equal(suite.T(), "failed to create employee", actual)
}

func (suite *WarehouseTeamRepoTestSuite) TestWarehouseTeamUpdate_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "phone", "photo"}).AddRow(dummyAdmin[0].ID, dummyAdmin[0].Name, dummyAdmin[0].Email, dummyAdmin[0].Password, dummyAdmin[0].Phone, dummyAdmin[0].Photo)

	suite.mockSql.ExpectQuery("SELECT id, name, email, password, phone, photo FROM admin_wh WHERE email = ?").WithArgs(dummyAdmin[0].Email).WillReturnRows(rows)

	suite.mockSql.ExpectExec("UPDATE admin_wh SET id = \\$1, name = \\$2, email = \\$3, password = \\$4, phone = \\$5, photo = \\$6 WHERE email = \\$7").
		WithArgs(dummyAdmin[0].ID, dummyAdmin[0].Name, dummyAdmin[0].Email, "67890", dummyAdmin[0].Phone, dummyAdmin[0].Photo, dummyAdmin[0].Email).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewWarehouseTeamRepo(suite.mockDb)
	updatedEmployee := &entity.WarehouseTeam{
		ID:       dummyAdmin[0].ID,
		Name:     dummyAdmin[0].Name,
		Email:    dummyAdmin[0].Email,
		Password: "67890",
		Phone:    dummyAdmin[0].Phone,
		Photo:    dummyAdmin[0].Photo,
	}
	expected := fmt.Sprintf("employee %s updated successfully", updatedEmployee.Name)
	actual := repo.Update(updatedEmployee)
	if actual != expected {
		_, err := repo.GetByEmail(dummyAdmin[0].Email)
		fmt.Println("Update error:", err)
	}

	assert.Equal(suite.T(), expected, actual)
}

func (suite *WarehouseTeamRepoTestSuite) TestWarehouseTeamUpdate_Failed() {
	suite.mockSql.ExpectQuery("SELECT id, name, email, password, phone, photo FROM admin_wh WHERE email = ?").WithArgs(dummyAdmin[0].Email).WillReturnError(errors.New("employee not found"))

	repo := NewWarehouseTeamRepo(suite.mockDb)
	expected := "employee not found"
	actual := repo.Update(&dummyAdmin[0])

	assert.Equal(suite.T(), expected, actual)
}

func (suite *WarehouseTeamRepoTestSuite) TestWarehouseTeamDelete_Success() {
	suite.mockSql.ExpectQuery("SELECT photo FROM admin_wh WHERE id = ?").WithArgs(dummyAdmin[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(dummyAdmin[0].ID))

	suite.mockSql.ExpectExec("DELETE FROM admin_wh WHERE id = ?").WithArgs(dummyAdmin[0].ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewWarehouseTeamRepo(suite.mockDb)
	expected := fmt.Sprintf("employee with id %s deleted successfully", dummyAdmin[0].ID)
	actual := repo.Delete(dummyAdmin[0].ID)
	if actual != expected {
		err := repo.GetById(dummyAdmin[0].ID)
		fmt.Println("Update error:", err)
	}

	assert.Equal(suite.T(), expected, actual)
}

func (suite *WarehouseTeamRepoTestSuite) TestWarehouseTeamDelete_Failed() {
	suite.mockSql.ExpectQuery("SELECT photo FROM admin_wh WHERE id = ?").WithArgs(dummyAdmin[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}))

	repo := NewWarehouseTeamRepo(suite.mockDb)
	expected := "employee not found"
	actual := repo.Delete(dummyAdmin[0].ID)

	assert.Equal(suite.T(), expected, actual)

}

func (suite *WarehouseTeamRepoTestSuite) TestWarehouseTeamDelete_NotFound() {
	suite.mockSql.ExpectQuery("SELECT photo FROM admin_wh WHERE id = ?").WithArgs(dummyAdmin[0].ID).WillReturnRows(sqlmock.NewRows([]string{"id"}))

	repo := NewWarehouseTeamRepo(suite.mockDb)
	expected := "employee not found"
	actual := repo.Delete(dummyAdmin[0].ID)

	assert.Equal(suite.T(), expected, actual)
}

func (suite *WarehouseTeamRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatalln("error when opening a stub database connection", err)
	}
	suite.mockDb = mockDb
	suite.mockSql = mockSql
}

func (suite *WarehouseTeamRepoTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func TestWarehouseTeamRepoTestSuite(t *testing.T) {
	suite.Run(t, new(WarehouseTeamRepoTestSuite))
}
