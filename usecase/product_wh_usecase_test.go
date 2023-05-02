package usecase

import (
	"go_inven_ctrl/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyProducts = []entity.ProductWh{
	{
		ID:              "1",
		ProductName:     "Oreo",
		Price:           3000,
		ProductCategory: "food",
		Stock:           1000,
	},
	{
		ID:              "2",
		ProductName:     "Beng-beng",
		Price:           2500,
		ProductCategory: "food",
		Stock:           1000,
	},
	{
		ID:              "3",
		ProductName:     "Indomie",
		Price:           3500,
		ProductCategory: "food",
		Stock:           1000,
	},
}

type repoMockProductWh struct {
	mock.Mock
}

type ProductWhUsecaseTestSuite struct {
	repoMockProductWh *repoMockProductWh
	suite.Suite
}

func (r *repoMockProductWh) GetAll() any {
	arg := r.Called()
	if arg.Get(0) == nil {
		return []entity.ProductWh{}
	}
	return arg.Get(0).([]entity.ProductWh)
}

func (r *repoMockProductWh) GetById(id string) any {
	arg := r.Called(id)
	if arg.Get(0) == nil {
		return "product not found"
	}
	return arg.Get(0)
}

func (r *repoMockProductWh) GetByName(name string) (*entity.ProductWh, error) {
	arg := r.Called(0)
	if arg[0] != nil {
		return &entity.ProductWh{}, arg.Error(0)
	}
	return &entity.ProductWh{}, nil
}

func (r *repoMockProductWh) Create(newProduct *entity.ProductWh) (*entity.ProductWh, error) {
	arg := r.Called(newProduct)
	if arg[0] != nil {
		return &entity.ProductWh{}, arg.Error(0)
	}
	return newProduct, nil
}

func (r *repoMockProductWh) Update(product *entity.ProductWh) string {
	arg := r.Called(product)
	if arg[0] != nil {
		return arg.String(0)
	}
	return "product updated"
}

func (r *repoMockProductWh) Delete(id string) string {
	arg := r.Called(id)
	if arg[0] != nil {
		return arg.String(0)
	}
	return "product deleted"
}

func (suite *ProductWhUsecaseTestSuite) TestInputProduct_Success() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("Create", &dummyProducts[0]).Return("product created")

	productWh, err := productWhUc.Input(&dummyProducts[0])

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "product created", productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestInputProduct_Failed() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("Create", &dummyProducts[0]).Return("failed to create product")

	productWh, err := productWhUc.Input(&dummyProducts[0])

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed to create product", productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestFindProducts_Success() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("GetAll").Return(dummyProducts)

	productWh := productWhUc.FindProducts()
	productWhs := productWh.([]entity.ProductWh)

	assert.Equal(suite.T(), dummyProducts, productWh)
	assert.Equal(suite.T(), len(dummyProducts), len(productWhs))
}

func (suite *ProductWhUsecaseTestSuite) TestFindProducts_Failed() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("GetAll").Return([]entity.ProductWh{})

	productWh := productWhUc.FindProducts()
	productWhs := productWh.([]entity.ProductWh)

	assert.Equal(suite.T(), 0, len(productWhs))
	assert.Empty(suite.T(), productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestFindProductById_Success() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("GetById", "1").Return(dummyProducts[0].ID)

	productWh := productWhUc.FindProductById("1")

	assert.Equal(suite.T(), dummyProducts[0].ID, productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestFindProductById_Failed() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("GetById", "5").Return("no data")

	productWh := productWhUc.FindProductById("5")

	assert.Equal(suite.T(), "no data", productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestFindProductByName_Success() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("GetByName", "1").Return(dummyProducts[0])

	productWh, err := productWhUc.FindProductByName("Oreo")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyProducts[0], productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestFindProductByName_Failed() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("GetByName", "Yakult").Return("no data")

	productWh, err := productWhUc.FindProductByName("Yakult")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "no data", productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestEditEmployee_Success() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("Update", &dummyProducts[0]).Return("product updated")

	productWh := productWhUc.Edit(&dummyProducts[0])

	assert.Equal(suite.T(), "product updated", productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestEditEmployee_Failed() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("Update", &dummyProducts[0]).Return("failed to update product")

	productWh := productWhUc.Edit(&dummyProducts[0])

	assert.Equal(suite.T(), "failed to update product", productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestUnregEmployee_Success() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("Delete", "1").Return("product deleted")

	productWh := productWhUc.Output("1")

	assert.Equal(suite.T(), "product deleted", productWh)
}

func (suite *ProductWhUsecaseTestSuite) TestUnregEmployee_Failed() {
	productWhUc := NewProductWhUsecase(suite.repoMockProductWh)
	suite.repoMockProductWh.On("Delete", "5").Return("failed to delete product")

	productWh := productWhUc.Output("5")

	assert.Equal(suite.T(), "failed to delete product", productWh)
}

func (suite *ProductWhUsecaseTestSuite) SetupTest() {
	suite.repoMockProductWh = new(repoMockProductWh)
}

func TestProductWhUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(ProductWhUsecaseTestSuite))
}
