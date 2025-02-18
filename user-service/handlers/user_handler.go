package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Murodkadirkhanoff/pet-microservice-golang-app/db"
	"github.com/Murodkadirkhanoff/pet-microservice-golang-app/dto"
	entity "github.com/Murodkadirkhanoff/pet-microservice-golang-app/models"
	pb "github.com/Murodkadirkhanoff/pet-microservice-golang-app/proto/auth"
	"github.com/Murodkadirkhanoff/pet-microservice-golang-app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

type ValidationErrorResponse struct {
	Field   string
	Message string
}

var validate *validator.Validate

func (s *AuthServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	userDTO := &dto.RegisterDTO{
		Email:    request.Email,
		Password: request.Password,
	}

	var validate *validator.Validate
	validate = validator.New()

	// Валидация
	err := validate.Struct(userDTO)

	if err != nil {
		// Если произошла ошибка валидации
		var validationErrors []ValidationErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			validationError := ValidationErrorResponse{
				Field:   err.Field(),
				Message: fmt.Sprintf("Field '%s' is invalid: %s", err.Field(), err.Tag()),
			}
			validationErrors = append(validationErrors, validationError)
		}

		// Формируем строку с ошибками для клиента
		validationErrorMessages := fmt.Sprintf("Validation failed: %v", validationErrors)

		// Возвращаем ошибку с детальной информацией
		return nil, status.Errorf(codes.InvalidArgument, validationErrorMessages)
	}

	if err := db.DB.Where("email = ?", userDTO.Email).First(&entity.User{}).Error; err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "User with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(userDTO.Password)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Password hashing failed")
	}

	user := entity.User{
		ID:       uuid.New(), // генерируем UUID вручную
		Email:    userDTO.Email,
		Password: hashedPassword,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Failed to create user: %v", err)
	}

	// Возвращаем ответ
	return &pb.RegisterResponse{
		Status:  int32(codes.OK),
		Message: "User registered successfully",
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	if request.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email is required")
	}
	if request.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Password is required")
	}

	loginDTO := dto.LoginDTO{
		Email:    request.Email,
		Password: request.Password,
	}

	userId, err := entity.ValidateCredentials(db.DB, loginDTO)

	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	token, err := utils.GenerateToken(loginDTO.Email, userId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not generate token")
	}

	userToken := pb.UserToken{
		UserId: userId.String(),
		Token:  token,
	}

	return &pb.LoginResponse{
		UserToken: &userToken,
	}, nil
}

func (s *AuthServer) Profile(ctx context.Context, request *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	userId, err := getUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user := entity.User{}

	result := db.DB.Take(&user, "id = ?", userId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve user: %v", result.Error)
	}

	userData := &pb.User{
		Id:         user.ID.String(),
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		IsActive:   user.IsActive,
		IsVerified: user.IsVerified,
		Avatar:     user.Avatar,
		CreatedAt:  user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.Format(time.RFC3339),
	}

	return &pb.ProfileResponse{
		Status:  int32(codes.OK),
		Message: "Success",
		Data: map[string]*pb.User{
			"user": userData,
		},
	}, nil
}

func (s *AuthServer) VerifyToken(ctx context.Context, request *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	_, err := getUserId(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Authentication failed: %v", err)
	}

	return &pb.VerifyTokenResponse{
		Status:  int32(codes.OK),
		Message: "success",
	}, nil
}

func (s *AuthServer) SaveProfile(ctx context.Context, request *pb.SaveProfileRequest) (*pb.SaveProfileResponse, error) {

	if request.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email is required")
	}

	userId, err := getUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := entity.User{}

	result := db.DB.Take(&user, "id = ?", userId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	user.Email = request.Email
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	db.DB.Save(&user)

	userData := &pb.User{
		Id:         user.ID.String(),
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		IsActive:   user.IsActive,
		IsVerified: user.IsVerified,
		Avatar:     user.Avatar,
		CreatedAt:  user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.Format(time.RFC3339),
	}

	return &pb.SaveProfileResponse{
		Status:  int32(codes.OK),
		Message: "success",
		Data: map[string]*pb.User{
			"user": userData,
		},
	}, nil
}

func getUserId(ctx context.Context) (string, error) {
	// Извлечение метаданных из контекста
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("missing metadata in context")
	}

	// Проверка наличия заголовка "authorization"
	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return "", errors.New("missing authorization header")
	}

	// Извлечение токена
	jwtToken := authHeader[0]
	if jwtToken == "" {
		return "", errors.New("authorization token is empty")
	}

	// Проверка токена
	userId, err := utils.VerifyToken(jwtToken)
	if err != nil {
		return "", fmt.Errorf("token verification failed: %v", err)
	}

	return userId, nil
}
