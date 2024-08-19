package database_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

type DatabaseTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *DatabaseTestSuite) SetupSuite() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		suite.T().Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		suite.T().Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	suite.Require().NoError(err, "Error connecting to the test database")

	suite.db = db.Debug()

	err = suite.db.AutoMigrate(&User{})
	suite.Require().NoError(err, "Error auto-migrating database tables")
}

func (suite *DatabaseTestSuite) TestUserInsertion() {
	user := User{Name: "John Doe"}
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err, "Error creating user record")

	var retrievedUser User
	err = suite.db.First(&retrievedUser, "name = ?", "John Doe").Error
	suite.Require().NoError(err, "Error retrieving user record")

	suite.Equal(user.Name, retrievedUser.Name, "Names should match")
}

func (suite *DatabaseTestSuite) TearDownSuite() {
	err := suite.db.Exec("DROP TABLE users;").Error
	suite.Require().NoError(err, "Error dropping test table")

	sqlDB, err := suite.db.DB()
	suite.Require().NoError(err, "Error getting SQL connection from GORM")

	err = sqlDB.Close()
	suite.Require().NoError(err, "Error closing test database connection")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
