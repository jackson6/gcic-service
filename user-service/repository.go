// user-service/repository.go

package main

import (
	pb "github.com/jackson6/gcic-service/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(*pb.User) error
	All() ([]*pb.User, error)
	Get(string) (*pb.User, error)
	GetByEmail(string) (*pb.User, error)
	GetUsers([]string) ([]*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

// Create a new user
func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetAll users
func (repo *UserRepository) All() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUsers(id []string) ([]*pb.User, error) {
	users := []*pb.User{}
	if err := repo.db.Where("id IN (?)", id).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//func (repo *UserRepository) GetUsers(id []string) ([]*pb.User, error) {
//	var users []*pb.User
//	rows, err := repo.db.Raw(`select id, first_name, last_name, dob, email, trn, parish, country, address, cell_phone,
//									home_phone, address2, member_id, gender, initial, plan_id, sponsor_id, referral_code
//									from users where IN (?)`, id).Rows()
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		user := new(pb.User)
//		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Dob, &user.Email, &user.Trn, &user.Parish,
//					&user.Country, &user.Address, &user.CellPhone, &user.HomePhone, &user.Address2, &user.MemberId,
//					&user.Gender, &user.Initial, &user.PlanId, &user.SponsorId, &user.ReferralCode)
//		users = append(users, user)
//	}
//
//	return users, nil
//}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
