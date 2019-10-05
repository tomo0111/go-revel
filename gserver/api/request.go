package api

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/go-playground/validator.v9"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/cache"
	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var rhInstance Request

type Request interface {
	// Http interceptor
	// Verify http header
	// If it has request body, verify bind request body
	Intercept(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResBody)

	// If need to verify authentication, verify authentication and authorization
	VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResBody)

	// Validate http request
	ValidateBody(w http.ResponseWriter, i interface{}) *model.ErrorResBody

	// Validate http header
	validateHeader(r *http.Request) *model.ErrorResBody

	// Bind http request body
	bindBody(r *http.Request) ([]byte, *model.ErrorResBody)

	verifyOperatorMember(token string) (*model.AuthUser, *model.ErrorResBody)

	verifyServiceMember(token string) (*model.AuthUser, *model.ErrorResBody)

	verifyToken(token string) (*model.AuthUser, *model.ErrorResBody)
}

type RequestImpl struct {
	userService           service.UserService
	operatorPolicyService service.OperatorPolicyService
	redisClient           cache.RedisClient
}

func GetRequestInstance() Request {
	if rhInstance == nil {
		rhInstance = NewRequest()
	}
	return rhInstance
}

func NewRequest() Request {
	log.Logger.Info("New `Request` instance")
	log.Logger.Info("Inject `AuthService` to `Request`")
	return RequestImpl{
		userService:           service.GetUserServiceInstance(),
		operatorPolicyService: service.GetOperatorPolicyServiceInstance(),
		redisClient:           cache.GetRedisClientInstance(),
	}
}

func (rh RequestImpl) Intercept(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResBody) {
	if err := rh.validateHeader(r); err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	// Set request scope
	apiKey := r.Header.Get("Api-Key")
	token := r.Header.Get("Authorization")
	ctx.SetApiKey(apiKey)
	ctx.SetToken(token)

	bodyBytes, err := rh.bindBody(r)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return nil, err
	}

	return bodyBytes, nil
}

func (rh RequestImpl) ValidateBody(w http.ResponseWriter, i interface{}) *model.ErrorResBody {
	err := validator.New().Struct(i)
	if err != nil {
		log.Logger.Info("Request is invalid")
		errModel := model.BadRequest("Failed to request validation.")
		model.Error(w, errModel.ToJson(), errModel.Code)
		return errModel
	}
	return nil
}

func (rh RequestImpl) VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResBody) {
	var authUser *model.AuthUser
	var err *model.ErrorResBody
	switch authType {
	case property.AuthOperator:
		authUser, err = rh.verifyOperatorMember(ctx.GetToken().(string))
	case property.AuthUser:
		authUser, err = rh.verifyServiceMember(ctx.GetToken().(string))
	}

	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return nil, err
	}
	return authUser, err
}

func (rh RequestImpl) validateHeader(r *http.Request) *model.ErrorResBody {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Logger.Info("Not allowed content-type")
		return model.BadRequest("Need to content type is only json.")
	}
	return nil
}

func (rh RequestImpl) bindBody(r *http.Request) ([]byte, *model.ErrorResBody) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, model.InternalServerError("Error request body bind")
	}

	return body, nil
}

func (rh RequestImpl) verifyOperatorMember(token string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := rh.verifyToken(token)
	if err != nil {
		return nil, err
	}

	// TODO: Read cache

	operatorRole, err := rh.operatorPolicyService.GetByUserIdAndRoleId(authUser.UserId, authUser.RoleId)
	if err != nil {
		return nil, err
	}

	if operatorRole == nil {
		log.Logger.Info("OperatorMemberRole data is null")
		return nil, model.Unauthorized("Failed to token.")
	}

	return authUser, nil
}

func (rh RequestImpl) verifyServiceMember(token string) (*model.AuthUser, *model.ErrorResBody) {
	authUser, err := rh.verifyToken(token)
	if err != nil {
		return nil, err
	}

	return authUser, nil
}

func (rh RequestImpl) verifyToken(token string) (*model.AuthUser, *model.ErrorResBody) {
	if !strings.Contains(token, "Bearer") {
		log.Logger.Info("Not found authorization header or not contain `Bearer` in authorization header")
		return nil, model.Unauthorized("Unauthorized.")
	}

	jwt := strings.Replace(token, "Bearer ", "", 1)
	userData, result := rh.userService.ParseJwt(jwt)
	if !result {
		log.Logger.Info("Failed to parse token")
		return nil, model.Unauthorized("Failed to token.")
	}

	// TODO: Read cache
	id, _ := strconv.Atoi(userData["user_id"])
	user, err := rh.userService.GetUserById(id)
	if err != nil {
		return nil, model.Unauthorized("Failed to token.")
	}

	if user == nil {
		log.Logger.Info("User data is null")
		return nil, model.Unauthorized("Failed to token.")
	}

	roleId, _ := strconv.Atoi(userData["role"])
	uid, _ := uuid.FromString(userData["user_uuid"])
	authUser := model.AuthUser{
		Username: userData["user_name"],
		UserUuid: uid,
		UserId:   id,
		Expires:  userData["expires"],
		RoleId:   roleId,
	}
	return &authUser, nil
}
