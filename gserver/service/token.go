package service

import (
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/config"
	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var tsInstance TokenService

type TokenService interface {
	// Generate jwt token
	Generate(userType string, groupIdStr string, userEntity entity.User) (*string, *model.ErrorResBody)

	// Parse and check token
	ParseToken(token string) (map[string]string, bool)

	// Verify operator token
	VerifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody)

	// Verify user token
	VerifyUserToken(token string) (*model.AuthUser, *model.ErrorResBody)

	// Generate signed in token
	signedInToken(userId int, userUuid string, roleId int, serviceId int, policyId int) *string

	// Generate operator token
	generateOperatorToken(userEntity entity.User) (*string, *model.ErrorResBody)

	// Generate user token
	generateUserToken(userEntity entity.User, groupId int) (*string, *model.ErrorResBody)

	// Get auth user data in token
	getAuthUserInToken(token string) (*model.AuthUser, *model.ErrorResBody)
}

// TokenService struct
type tokenServiceImpl struct {
	userService           UserService
	operatorPolicyService OperatorPolicyService
	userServiceService    UserServiceService
	service               Service
	policyService         PolicyService
	appConfig             config.AppConfig
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetTokenServiceInstance() TokenService {
	if tsInstance == nil {
		tsInstance = NewTokenService()
	}
	return tsInstance
}

// Constructor
func NewTokenService() TokenService {
	log.Logger.Info("New `TokenService` instance")
	log.Logger.Info("Inject `UserGroup`, `OperatorPolicyService` to `TokenService`")
	return tokenServiceImpl{
		userService:           GetUserServiceInstance(),
		operatorPolicyService: GetOperatorPolicyServiceInstance(),
		userServiceService:    GetUserServiceServiceInstance(),
		service:               GetServiceInstance(),
		policyService:         GetPolicyServiceInstance(),
		appConfig:             config.App,
	}
}

func (tsi tokenServiceImpl) Generate(userType string, groupIdStr string, userEntity entity.User) (*string, *model.ErrorResBody) {
	if strings.EqualFold(groupIdStr, "") {
		groupIdStr = "0"
	}

	groupId, err := strconv.Atoi(groupIdStr)
	if err != nil {
		return nil, model.BadRequest("Group id is only integer of query parameter")
	}

	switch userType {
	case property.AuthOperator:
		return tsi.generateOperatorToken(userEntity)
	case property.AuthUser:
		return tsi.generateUserToken(userEntity, groupId)
	case "":
		return tsi.generateUserToken(userEntity, groupId)
	default:
		return nil, model.BadRequest("Not support type of query parameter")
	}
}

func (tsi tokenServiceImpl) ParseToken(token string) (map[string]string, bool) {
	resultMap := map[string]string{}

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(tsi.appConfig.PrivateKeyBase64), nil
	})

	if err != nil || !parseToken.Valid {
		log.Logger.Error("Failed to parse token validation", err.Error())
		return resultMap, false
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	if _, ok := claims["user_uuid"].(string); !ok {
		log.Logger.Info("Can not get user_uuid from token")
		return resultMap, false
	}
	if _, ok := claims["user_id"].(string); !ok {
		log.Logger.Info("Can not get user_id from token")
		return resultMap, false
	}
	if _, ok := claims["expires"].(string); !ok {
		log.Logger.Info("Can not get expires from token")
		return resultMap, false
	}
	if _, ok := claims["role_id"].(string); !ok {
		log.Logger.Info("Can not get role_id from token")
		return resultMap, false
	}
	if _, ok := claims["service_id"].(string); !ok {
		log.Logger.Info("Can not get service_id from token")
		return resultMap, false
	}

	resultMap["user_uuid"] = claims["user_uuid"].(string)
	resultMap["user_id"] = claims["user_id"].(string)
	resultMap["expires"] = claims["expires"].(string)
	resultMap["role_id"] = claims["role_id"].(string)
	resultMap["service_id"] = claims["service_id"].(string)

	return resultMap, true
}

func (tsi tokenServiceImpl) VerifyOperatorToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := tsi.getAuthUserInToken(token)
	if err != nil {
		return nil, err
	}

	operatorRole, err := tsi.operatorPolicyService.GetByUserIdAndRoleId(authUser.UserId, authUser.RoleId)
	if operatorRole == nil || err != nil {
		log.Logger.Info("Not contain operator role or failed to query")
		return nil, model.Forbidden("Forbidden this token")
	}

	return authUser, nil
}

func (tsi tokenServiceImpl) VerifyUserToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := tsi.getAuthUserInToken(token)
	if err != nil {
		return nil, err
	}

	userService, err := tsi.userServiceService.GetUserServiceByUserIdAndServiceId(authUser.UserId, authUser.ServiceId)
	if userService == nil || err != nil {
		log.Logger.Info("Not contain service of user or failed to query")
		return nil, model.Forbidden("Forbidden this token")
	}

	return authUser, nil
}

func (tsi tokenServiceImpl) generateOperatorToken(userEntity entity.User) (*string, *model.ErrorResBody) {
	// TODO: Cache
	targetUser, err := tsi.userService.GetUserWithOperatorPolicyByEmail(userEntity.Email)
	if err != nil || targetUser == nil {
		return nil, model.BadRequest("Failed to email or password")
	}

	if !tsi.userService.ComparePw(targetUser.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}

	if targetUser.OperatorPolicy.RoleId != property.OperatorRoleId {
		return nil, model.BadRequest("Can not issue token")
	}

	// Operator token is not required service id, policy id
	serviceId := 0
	policyId := 0
	return tsi.signedInToken(targetUser.UserId, targetUser.Uuid.String(), targetUser.OperatorPolicy.RoleId, serviceId, policyId), nil
}

func (tsi tokenServiceImpl) generateUserToken(userEntity entity.User, groupId int) (*string, *model.ErrorResBody) {
	// TODO: Cache service data
	service, err := tsi.service.GetServiceByApiKey(ctx.GetApiKey().(string))
	if err != nil || service == nil {
		return nil, model.BadRequest("Not found registered services by Api-Key")
	}

	// TODO: Cache user data
	targetUser, err := tsi.userService.GetUserByEmail(userEntity.Email)
	if err != nil || targetUser == nil {
		return nil, model.BadRequest("Failed to email or password")
	}

	if !tsi.userService.ComparePw(targetUser.Password, userEntity.Password) {
		return nil, model.BadRequest("Failed to email or password")
	}

	// Case of group_id is zero, not using policy.
	if groupId == 0 {
		roleId := 0
		policyId := 0
		return tsi.signedInToken(targetUser.Id, targetUser.Uuid.String(), roleId, service.Id, policyId), nil
	}

	policy, err := tsi.policyService.GetPolicyByUserGroup(targetUser.Id, groupId)
	if err != nil {
		return nil, err
	}

	return tsi.signedInToken(targetUser.Id, targetUser.Uuid.String(), 0, service.Id, policy.Id), nil
}

func (tsi tokenServiceImpl) signedInToken(userId int, userUuid string, roleId int, serviceId int, policyId int) *string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = strconv.Itoa(userId)
	claims["user_uuid"] = userUuid
	claims["expires"] = time.Now().Add(time.Hour * 1).String()
	claims["role_id"] = strconv.Itoa(roleId)
	claims["service_id"] = strconv.Itoa(serviceId)
	claims["policy_id"] = strconv.Itoa(policyId)

	signedToken, err := token.SignedString([]byte(tsi.appConfig.PrivateKeyBase64))
	if err != nil {
		log.Logger.Error("Failed to issue signed token", err.Error())
		return nil
	}

	return &signedToken
}

func (tsi tokenServiceImpl) getAuthUserInToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	if !strings.Contains(token, "Bearer") {
		log.Logger.Info("Not found authorization header or not contain `Bearer` in authorization header")
		return nil, model.Unauthorized("Unauthorized.")
	}

	userData, result := tsi.ParseToken(strings.Replace(token, "Bearer ", "", 1))
	if !result {
		return nil, model.Unauthorized("Failed to token.")
	}

	userId, _ := strconv.Atoi(userData["user_id"])
	userUuid, _ := uuid.Parse(userData["user_uuid"])
	roleId, _ := strconv.Atoi(userData["role_id"])
	serviceId, _ := strconv.Atoi(userData["service_id"])
	policyId, _ := strconv.Atoi(userData["policy_id"])

	authUser := &model.AuthUser{
		UserId:    userId,
		UserUuid:  userUuid,
		ServiceId: serviceId,
		Expires:   userData["expires"],
		RoleId:    roleId,
		PolicyId:  policyId,
	}

	return authUser, nil
}
