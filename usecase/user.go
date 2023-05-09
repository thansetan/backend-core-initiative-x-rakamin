package usecase

import (
	"context"
	"errors"
	"fmt"
	"self-payroll/model"
	"self-payroll/request"
	"self-payroll/response"
	"self-payroll/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type userUsecase struct {
	userRepository  model.UserRepository
	positionRepo    model.PositionRepository
	companyRepo     model.CompanyRepository
	transactionRepo model.TransactionRepository
}

func NewUserUsecase(user model.UserRepository, post model.PositionRepository, transactionRepo model.TransactionRepository, company model.CompanyRepository) model.UserUsecase {
	return &userUsecase{userRepository: user, positionRepo: post, transactionRepo: transactionRepo, companyRepo: company}
}

func (p *userUsecase) WithdrawSalary(ctx context.Context, userID int, req *request.WithdrawRequest) error {
	user, err := p.userRepository.FindByID(ctx, userID)
	fmt.Println("start")
	if err != nil {
		fmt.Println("error di cari ID")
		return err
	}

	if user.SecretID != req.SecretID {
		return errors.New("secret id not valid")
	}

	err = p.transactionRepo.CheckLastWithdrawal(ctx, userID)
	if err != nil {
		fmt.Println("error di cek last withdrawal")
		return err
	}
	notes := fmt.Sprintf("%s (%d) withdraw salary", user.Name, user.ID)

	err = p.companyRepo.DebitBalance(ctx, user.Position.Salary, userID, notes)
	if err != nil {
		fmt.Println("error di debit balance")
		return err
	}
	return nil
}

func (p *userUsecase) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := p.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userUsecase) FetchUser(ctx context.Context, limit, offset int) ([]*model.User, error) {

	users, err := p.userRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (p *userUsecase) DestroyUser(ctx context.Context, id int) error {
	err := p.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *userUsecase) EditUser(ctx context.Context, id int, req *request.UpdateRequest) (*model.User, error) {
	_, err := p.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user, err := p.userRepository.UpdateByID(ctx, id, &model.User{
		SecretID: req.SecretID,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userUsecase) StoreUser(ctx context.Context, req *request.RegisterRequest) (*model.User, error) {
	hashedPass, hashErr := utils.GeneratePassword(req.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	newUser := &model.User{
		SecretID:   req.SecretID,
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashedPass,
		Phone:      req.Phone,
		Address:    req.Address,
		PositionID: req.PositionID,
	}

	_, err := p.positionRepo.FindByID(ctx, req.PositionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("position id not valid ")
		}

		return nil, err
	}

	user, err := p.userRepository.Create(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userUsecase) AdminRegister(ctx context.Context, req *request.AdminRequest) (*model.User, error) {
	hashedPass, hashErr := utils.GeneratePassword(req.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	newUser := &model.User{
		SecretID:   req.SecretID,
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashedPass,
		Phone:      req.Phone,
		Address:    req.Address,
		IsAdmin:    true,
		PositionID: 69420,
	}

	user, err := p.userRepository.Create(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userUsecase) Login(ctx context.Context, req *request.LoginRequest) (*response.LoginResponse, error) {
	userData, userErr := p.userRepository.FindByEmail(ctx, req.Email)
	if userErr != nil {
		return nil, userErr
	}
	if !utils.ValidatePassword(req.Password, userData.Password) {
		return nil, fmt.Errorf("invalid email/password")
	}
	claims := jwt.MapClaims{
		"id":         userData.ID,
		"positionID": userData.PositionID,
		"isAdmin":    userData.IsAdmin,
		"exp":        time.Now().Add(time.Hour * 48).Unix(),
	}
	t, err := utils.GenerateJWT(claims)
	if err != nil {
		return nil, err
	}
	res := &response.LoginResponse{
		Token: t,
	}

	c, _ := utils.DecodeJWT(t)
	fmt.Println(c)
	return res, nil
}
