package database

import (
	"bonhart_auth_service/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/Gerardo115pp/patriots_lib/echo"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMYSQLrepository() (*MysqlRepository, error) {
	db, err := sql.Open("mysql", createDSN())
	if err != nil {
		return nil, err
	}
	// Load all customers from database

	return &MysqlRepository{db: db}, nil
}

func (mysql_repo *MysqlRepository) GetExpertByID(ctx context.Context, id int) (*models.Expert, error) {
	var target_expert *models.Expert = new(models.Expert)
	var err error

	err = mysql_repo.db.QueryRow("SELECT `id`, `username`,`email`, `password`, `is_active`, `is_available`, `type` FROM `experts` WHERE `id` = ?", id).Scan(&target_expert.ID, &target_expert.Username, &target_expert.Email, &target_expert.Password, &target_expert.IsActive, &target_expert.IsAvailable, &target_expert.ExpertType)

	echo.EchoDebug(fmt.Sprintf("Expert data:\n %+v", target_expert))
	return target_expert, err
}

func (mysql_repo *MysqlRepository) GetExpertByEmail(ctx context.Context, email string) (*models.Expert, error) {
	var targetExpert *models.Expert = new(models.Expert)
	var err error

	err = mysql_repo.db.QueryRow("SELECT `id`, `username`,`email`, `password`, `is_active`, `is_available`, `type` FROM `experts` WHERE `email` = ?", email).Scan(&targetExpert.ID, &targetExpert.Username, &targetExpert.Email, &targetExpert.Password, &targetExpert.IsActive, &targetExpert.IsAvailable, &targetExpert.ExpertType)
	return targetExpert, err
}

func (mysql_repo *MysqlRepository) GetAdminByID(ctx context.Context, id int) (*models.ItzAdmin, error) {
	var target_admin *models.ItzAdmin = new(models.ItzAdmin)
	var err error

	err = mysql_repo.db.QueryRow("SELECT `id`, `username`,`email`, `password`, `is_active`, `is_superadmin` FROM `admins` WHERE `id` = ?", id).Scan(&target_admin.ID, &target_admin.Username, &target_admin.Email, &target_admin.Password, &target_admin.IsActive, &target_admin.IsSuperadmin)

	return target_admin, err
}

func (mysql_repo *MysqlRepository) GetAdminByEmail(ctx context.Context, email string) (*models.ItzAdmin, error) {
	var target_admin *models.ItzAdmin = new(models.ItzAdmin)
	var err error

	err = mysql_repo.db.QueryRow("SELECT `id`, `username`,`email`, `password`, `is_active`, `is_superadmin` FROM `admins` WHERE `email` = ?", email).Scan(&target_admin.ID, &target_admin.Username, &target_admin.Email, &target_admin.Password, &target_admin.IsActive, &target_admin.IsSuperadmin)
	return target_admin, err
}

func (sqlite *MysqlRepository) Close() error {
	return sqlite.db.Close()
}
