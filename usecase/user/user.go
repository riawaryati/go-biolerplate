package user

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/furee/backend/domain/general"
	du "github.com/furee/backend/domain/user"
	"github.com/furee/backend/infra"
	"github.com/furee/backend/repo"
	ru "github.com/furee/backend/repo/user"
	"github.com/furee/backend/utils"
	"github.com/sirupsen/logrus"
)

type UserDataUsecaseItf interface {
	LoginUser(data du.UserLoginRequest) (string, error)
	VerifyOTP(data du.VerifyOTPRequest) (*general.JWTAccess, string, error)
}

type UserDataUsecase struct {
	Repo   ru.UserDataRepoItf
	DBList *infra.DatabaseList
	Conf   *general.SectionService
	Log    *logrus.Logger
}

func newUserDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) UserDataUsecase {
	return UserDataUsecase{
		Repo:   r.User.User,
		Conf:   conf,
		Log:    logger,
		DBList: dbList,
	}
}

func (uu UserDataUsecase) VerifyOTP(data du.VerifyOTPRequest) (*general.JWTAccess, string, error) {
	// phoneFilter := sha256.Sum256([]byte(data.Phone))
	// data.PhoneFilter = fmt.Sprintf("%x", phoneFilter[:])

	// isExist, err := uu.Repo.IsExistUser(data.PhoneFilter)
	// if err != nil {
	// 	uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
	// 	return nil, "", err
	// }

	// if !isExist {
	// 	uu.Log.WithField("request", utils.StructToString(data)).Errorf("user is not exist")
	// 	return nil, "Nomor Anda Belum Terdaftar", nil
	// }

	// user, err := uu.Repo.GetByPhone(data.PhoneFilter)
	// if err != nil {
	// 	uu.Log.WithField("user id", user.ID).WithError(err).Error("VerifyOTP | fail to get user data from repo")
	// 	return nil, "", err
	// }

	// if user == nil {
	// 	return nil, "", errors.New("user data not found")
	// }

	// // checking whitelist OTP. if phone is on whitelist, do not send OTP to this phone.
	// _, ok := uu.Conf.Whitelist.User.Phone[data.Phone]
	// if !ok {
	// 	if !user.OTP.Valid {
	// 		return nil, "", errors.New("OTP is null")
	// 	}

	// 	if user.OTP.String != data.OTP {
	// 		return nil, "Kode verifikasi salah. Silahkan cek kode verifikasi di akun whatsapp anda", errors.New("OTP not match")
	// 	}
	// } else {
	// 	if data.OTP != uu.Conf.Whitelist.User.OTP {
	// 		return nil, "Kode verifikasi salah. Silahkan cek kode verifikasi di akun whatsapp anda", errors.New("OTP not match")
	// 	}
	// }

	// err = uu.Repo.VerifyUser(nil, du.VerifyUser{PhoneFilter: data.PhoneFilter, OTP: user.OTP.String})
	// if err != nil {
	// 	return nil, "", err
	// }

	session, err := utils.GetEncrypt([]byte(uu.Conf.App.SecretKey), fmt.Sprintf("%v", 1))
	if err != nil {
		uu.Log.WithField("user id", 1).WithError(err).Error("VerifyOTP | fail to get token data from infra")
		return nil, "", err
	}

	accessToken, renewToken, err := utils.GenerateJWT(session)
	if err != nil {
		uu.Log.WithField("user id", 1).WithError(err).Error("VerifyOTP | fail to get token data from infra")
		return nil, "", err
	}

	generateTime := time.Now().UTC()

	return &general.JWTAccess{
		AccessToken:        accessToken,
		AccessTokenExpired: generateTime.Add(time.Duration(uu.Conf.Authorization.JWT.AccessTokenDuration) * time.Minute).Format(time.RFC3339),
		RenewToken:         renewToken,
		RenewTokenExpired:  generateTime.Add(time.Duration(uu.Conf.Authorization.JWT.RefreshTokenDuration*24) * time.Hour).Format(time.RFC3339),
	}, "success verify user account", nil
}

func (uu UserDataUsecase) LoginUser(data du.UserLoginRequest) (string, error) {
	phoneFilter := sha256.Sum256([]byte(data.Phone))
	phone := fmt.Sprintf("%x", phoneFilter[:])

	isExist, err := uu.Repo.IsExistUser(phone)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist user")
		return "", err
	}

	if !isExist {
		uu.Log.WithField("request", utils.StructToString(data)).Errorf("user is not exist")
		return "Nomor Anda Belum Terdaftar", errors.New("user not exist")
	}

	user, err := uu.Repo.GetByPhone(phone)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user data not found")
	}

	otpCode := utils.GenerateOTP()

	err = uu.Repo.UpdateOTP(nil, otpCode, user.ID)
	if err != nil {
		uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to update otp")
		return "fail to login user", nil
	}

	return "success send otp user", nil
}
