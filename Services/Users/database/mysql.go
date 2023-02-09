package database

import (
	"context"
	"database/sql"
	"fmt"
	"itz_customers_service/models"

	"github.com/Gerardo115pp/patriots_lib/echo"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
}

var local_cache map[int]*models.Expert /*  {ORDER_ID: ORDER_OBJECT}*/

func NewMYSQLrepository() (*MysqlRepository, error) {
	db, err := sql.Open("mysql", createDSN())
	if err != nil {
		return nil, err
	}

	new_sqlite_storage := MysqlRepository{db: db}
	new_sqlite_storage.loadCustomers() // Load all customers from database

	echo.Echo(echo.BlueBG, fmt.Sprintf("Loaded %d customers from database", len(local_cache)))

	return &MysqlRepository{db: db}, nil
}

func (mysql_repo *MysqlRepository) InsertExpert(ctx context.Context, expert *models.Expert) error {
	if _, exists := local_cache[expert.ID]; exists || expert.ID != 0 {
		return fmt.Errorf("Customer %d already exists", expert.ID)
	}

	stmt, err := mysql_repo.db.Prepare("INSERT INTO experts (username, name, email, password, type, created_by) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, expert.Username, expert.Name, expert.Email, expert.Password, expert.ExpertType, expert.CreatedBy)
	if err != nil {
		return err
	}

	expert, err = mysql_repo.GetExpertByEmail(ctx, expert.Email)

	local_cache[expert.ID] = expert
	return nil
}

func (mysql_repo *MysqlRepository) GetExpertByID(ctx context.Context, id int) (*models.Expert, error) {
	var target_expert *models.Expert = new(models.Expert)
	var err error

	if targetExpert, exists := local_cache[id]; exists {
		return targetExpert, nil
	}

	echo.EchoWarn(fmt.Sprintf("Expert %d not found in cache", id))

	err = mysql_repo.db.QueryRow("SELECT id, username, name, email, password, is_active, is_available, created_at, last_login, type, created_by FROM experts WHERE id = ?", id).Scan(
		&target_expert.ID, &target_expert.Username, &target_expert.Name, &target_expert.Email, &target_expert.Password, &target_expert.IsActive, &target_expert.IsAvailable, &target_expert.CreatedAt, &target_expert.LastLogin, &target_expert.ExpertType, &target_expert.CreatedBy,
	)

	echo.EchoDebug(fmt.Sprintf("Expert data:\n %+v", target_expert))
	return target_expert, err
}

func (mysql_repo *MysqlRepository) GetExpertByEmail(ctx context.Context, email string) (*models.Expert, error) {
	var targetExpert *models.Expert = new(models.Expert)
	var err error

	err = mysql_repo.db.QueryRow("SELECT id, username, name, email, password, is_active, is_available, created_at, last_login, type, created_by FROM experts WHERE email = ?", email).Scan(
		&targetExpert.ID, &targetExpert.Username, &targetExpert.Name, &targetExpert.Email, &targetExpert.Password, &targetExpert.IsActive, &targetExpert.IsAvailable, &targetExpert.CreatedAt, &targetExpert.LastLogin, &targetExpert.ExpertType, &targetExpert.CreatedBy,
	)
	return targetExpert, err
}

func (mysql_repo *MysqlRepository) GetAllExperts(ctx context.Context) ([]*models.Expert, error) {
	var experts []*models.Expert

	rows, err := mysql_repo.db.Query("SELECT id, username, name, email, password, is_active, is_available, created_at, last_login, type, created_by FROM experts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var expert *models.Expert = new(models.Expert)

		err = rows.Scan(&expert.ID, &expert.Username, &expert.Name, &expert.Email, &expert.Password, &expert.IsActive, &expert.IsAvailable, &expert.CreatedAt, &expert.LastLogin, &expert.ExpertType, &expert.CreatedBy)
		if err != nil {
			return nil, err
		}

		experts = append(experts, expert)
	}

	return experts, nil
}

func (mysql_repo *MysqlRepository) UpdateExpert(ctx context.Context, expert *models.Expert) error {
	stmt, err := mysql_repo.db.Prepare("UPDATE `experts` SET `username`=?, `name`=?, `email`=?, `password`=?, `is_active`=?, `is_available`=? WHERE `id` = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, expert.Username, expert.Name, expert.Email, expert.Password, expert.IsActive, expert.IsAvailable, expert.ID)
	if err != nil {
		return err
	}

	return nil
}

func (mysql_repo *MysqlRepository) DeleteExpert(ctx context.Context, id int) error {
	_, err := mysql_repo.db.Exec("DELETE FROM experts WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (mysql_repo *MysqlRepository) GetExpertPublicProfile(ctx context.Context, expert_id int) (*models.PublicProfile, error) {
	var public_profile *models.PublicProfile = new(models.PublicProfile)
	var err error

	stmt, err := mysql_repo.db.Prepare("SELECT public_name, professional_title, description, brief, image_href FROM public_profiles WHERE expert=?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRowContext(ctx, expert_id).Scan(&public_profile.Public_name, &public_profile.Professional_title, &public_profile.Description, &public_profile.Brief, &public_profile.Image_url)
	if err != nil {
		return nil, err
	}

	public_profile.Expert_id = expert_id
	return public_profile, nil
}

func (mysql_repo *MysqlRepository) GetActiveExpertProfiles(ctx context.Context) ([]*models.PublicProfile, error) {
	var public_profiles []*models.PublicProfile

	rows, err := mysql_repo.db.QueryContext(ctx, "SELECT `pp`.`public_name`, `pp`.`professional_title`, `pp`.`description`, `pp`.`brief`, `pp`.`image_href`, `pp`.`expert`, `e`.`type` FROM `public_profiles` `pp` LEFT JOIN `experts` `e` ON `e`.`id`=`pp`.`expert` WHERE `e`.`is_active` = 1 AND `e`.`is_available` = 1")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var current_profile *models.PublicProfile = new(models.PublicProfile)
		err = rows.Scan(&current_profile.Public_name, &current_profile.Professional_title, &current_profile.Description, &current_profile.Brief, &current_profile.Image_url, &current_profile.Expert_id, &current_profile.ExpertType)
		if err != nil {
			return nil, err
		}

		public_profiles = append(public_profiles, current_profile)
	}

	return public_profiles, nil
}

func (mysql_repo *MysqlRepository) UpdateExpertPublicProfile(ctx context.Context, public_profile *models.PublicProfile) error {
	var err error
	var stmt *sql.Stmt

	stmt, err = mysql_repo.db.Prepare("UPDATE public_profiles SET public_name=?, professional_title=?, description=?, brief=?, image_href=? WHERE expert = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, public_profile.Public_name, public_profile.Professional_title, public_profile.Description, public_profile.Brief, public_profile.Image_url, public_profile.Expert_id)
	if err != nil {
		return err
	}

	return nil
}

func (sqlite *MysqlRepository) loadCustomers() error {
	if local_cache != nil {
		echo.EchoErr(fmt.Errorf("Experts are already loaded"))
	}
	echo.EchoDebug("Loading experts")

	local_cache = make(map[int]*models.Expert)

	rows, err := sqlite.db.Query("SELECT id, username, name, email, password, is_active, is_available, created_at, last_login, type, created_by FROM experts")
	if err != nil {
		echo.EchoErr(err)
		return err
	}

	for rows.Next() {
		var expert *models.Expert = new(models.Expert)
		err = rows.Scan(&expert.ID, &expert.Username, &expert.Name, &expert.Email, &expert.Password, &expert.IsActive, &expert.IsAvailable, &expert.CreatedAt, &expert.LastLogin, &expert.ExpertType, &expert.CreatedBy)
		if err != nil {
			echo.EchoErr(err)
			return err
		}

		local_cache[expert.ID] = expert
	}

	echo.EchoDebug(fmt.Sprintf("Loaded %d experts", len(local_cache)))

	return nil
}

func (sqlite *MysqlRepository) Close() error {
	return sqlite.db.Close()
}
