package session

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"hospital-system/server/app/services/session"
	"hospital-system/server/app/services/user"
	"net/http"
)

const (
	AuthorizationCookieName = "Authorization"

	Session_Login_BadRequestErr            = "bad_request"
	Session_Login_UserPasswordPairWrongErr = "user_password_pair_wrong"
	Session_Login_UserNotFoundErr          = "user_not_found"
	Session_Login_InactiveAccountErr       = "inactive_account"
	Session_Login_CreateSessionFailedErr   = "create_session_failed"

	Session_Logout_BadRequestErr          = "bad_request"
	Session_Logout_DeleteSessionFailedErr = "delete_session_failed"

	Session_Validate_BadRequestErr       = "bad_request"
	Session_Validate_SessionGetFailedErr = "session_get_failed"
	Session_Validate_SessionExpiredErr   = "session_expired"
	Session_Validate_UserNotFoundErr     = "user_not_found"
	Session_Validate_InactiveAccountErr  = "inactive_account"
)

type HandlerImpl struct {
	userService    user.Service
	sessionService session.Service
}

func NewHandler(userService user.Service, sessionService session.Service) *HandlerImpl {
	return &HandlerImpl{
		userService:    userService,
		sessionService: sessionService,
	}
}

func (this *HandlerImpl) Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Session_Login_BadRequestErr})
		return
	}

	existingUser, err := this.userService.GetByUsername(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": Session_Login_UserPasswordPairWrongErr})
		return
	}

	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": Session_Login_UserNotFoundErr})
		return
	} else if existingUser.Archived != nil && *existingUser.Archived {
		c.JSON(http.StatusForbidden, gin.H{"error": Session_Login_InactiveAccountErr})
		return
	}

	isPasswordValid, err := this.userService.ValidateUserPassword(existingUser.ID, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": Session_Login_UserNotFoundErr})
		return
	}
	if !isPasswordValid {
		c.JSON(http.StatusForbidden, gin.H{"error": Session_Login_UserPasswordPairWrongErr})
		return
	}

	jwtToken, err := this.sessionService.CreateSession(session.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: existingUser.ID.String(),
		},
		Username:    existingUser.Username,
		Firstname:   existingUser.Firstname,
		Lastname:    existingUser.Lastname,
		Role:        existingUser.Role,
		Team:        existingUser.Team,
		Permissions: existingUser.Permissions,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": Session_Login_CreateSessionFailedErr})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(AuthorizationCookieName, jwtToken, 900, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
		"user": user.User{
			Firstname:                    existingUser.Firstname,
			Lastname:                     existingUser.Lastname,
			NationalIdentificationNumber: existingUser.NationalIdentificationNumber,
			Username:                     existingUser.Username,
			Email:                        existingUser.Email,
			PhoneNumber:                  existingUser.PhoneNumber,
			MailingAddress:               existingUser.MailingAddress,
			City:                         existingUser.City,
			State:                        existingUser.State,
			Zip:                          existingUser.Zip,
			Gender:                       existingUser.Gender,
			Birthday:                     existingUser.Birthday,
			JoiningDate:                  existingUser.JoiningDate,
			Archived:                     existingUser.Archived,
			Role:                         existingUser.Role,
			Team:                         existingUser.Team,
			Permissions:                  existingUser.Permissions,
		},
	})
}

func (this *HandlerImpl) Logout(c *gin.Context) {
	var request LogoutRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Session_Logout_BadRequestErr})
		return
	}

	if err := this.sessionService.DeleteSession(request.Token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": Session_Logout_DeleteSessionFailedErr})
		return
	}

	c.SetCookie(AuthorizationCookieName, "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func (this *HandlerImpl) ValidateSession(c *gin.Context) {
	var request ValidateSessionRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Session_Validate_BadRequestErr})
		return
	}

	claims, err := this.sessionService.GetSession(request.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": Session_Validate_SessionGetFailedErr})
		c.Abort()
		return
	}
	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": Session_Validate_SessionExpiredErr})
		c.Abort()
		return
	}

	existingUser, err := this.userService.GetByUsername(claims.Username)
	if err != nil || existingUser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": Session_Validate_UserNotFoundErr})
		return
	} else if existingUser.Archived != nil && *existingUser.Archived {
		c.JSON(http.StatusForbidden, gin.H{"error": Session_Validate_InactiveAccountErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.User{
			Firstname:                    claims.Firstname,
			Lastname:                     claims.Lastname,
			NationalIdentificationNumber: existingUser.NationalIdentificationNumber,
			Username:                     existingUser.Username,
			Email:                        existingUser.Email,
			PhoneNumber:                  existingUser.PhoneNumber,
			MailingAddress:               existingUser.MailingAddress,
			City:                         existingUser.City,
			State:                        existingUser.State,
			Zip:                          existingUser.Zip,
			Gender:                       existingUser.Gender,
			Birthday:                     existingUser.Birthday,
			JoiningDate:                  existingUser.JoiningDate,
			Archived:                     existingUser.Archived,
			Role:                         claims.Role,
			Team:                         claims.Team,
			Permissions:                  claims.Permissions,
		},
	})
}
