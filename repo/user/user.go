package user

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	cg "github.com/furee/backend/constants/general"
	du "github.com/furee/backend/domain/user"
	"github.com/furee/backend/infra"
)

type UserDataRepo struct {
	DBList *infra.DatabaseList
}

func newUserDataRepo(dbList *infra.DatabaseList) UserDataRepo {
	return UserDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectUser = `
	SELECT
		user_id,
		name,
		status,
		phone,
		phone_filter,
		otp,
		otp_created_at,
		created_at,
		updated_at,
		updated_by
	FROM
		users`

	uqInsertUser = `
	INSERT INTO users (
		name,
		status,
		phone,
		phone_filter,
		otp,
		otp_created_at,
		created_at
	) VALUES (
		?, ?, ?, ?, ?, ?, ?
	)
	RETURNING user_id`

	uqUpdateUser = `
	UPDATE 
		users
	SET
		updated_by = ?,
		updated_at = NOW()`

	uqSelectExist = `
		SELECT EXISTS`

	uqWhere = `
	WHERE`

	uqFilterUserID = `
		user_id = ?`

	uqFilterLikeName = `
		lower(name) LIKE ?`

	uqFilterName = `
		lower(name) = ?`

	uqFilterPhone = `
		phone = ?`

	uqFilterPhoneFilter = `
		phone_filter = ?`

	uqFilterOTP = `
		otp = ?`

	uqFilterOTPCreatedAt = `
		otp_created_at = ?`

	uqFilterStatus = `
		status = ?`
)

type UserDataRepoItf interface {
	GetByID(userID int64) (*du.User, error)
	GetByName(userName string) ([]*du.User, error)
	GetByPhone(phoneFilter string) (*du.User, error)
	IsExistOTP(otp, phone string) (bool, error)
	IsExistUser(phone string) (bool, error)
	InsertUser(tx *sql.Tx, data du.CreateUser) (int64, error)
	VerifyUser(tx *sql.Tx, data du.VerifyUser) error
	UpdateStatus(tx *sql.Tx, status int, userID int64) error
	UpdateOTP(tx *sql.Tx, otp string, userID int64) error
}

func (ur UserDataRepo) GetByID(userID int64) (*du.User, error) {
	var res du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterUserID)
	query, args, err := ur.DBList.Backend.Read.In(q, userID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.ID == 0 {
		return nil, nil
	}

	return &res, nil
}

func (ur UserDataRepo) GetByName(userName string) ([]*du.User, error) {
	var res []*du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterLikeName)
	query, args, err := ur.DBList.Backend.Read.In(q, strings.ToLower(userName))
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Select(&res, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ur UserDataRepo) GetByPhone(phoneFilter string) (*du.User, error) {
	var res du.User

	q := fmt.Sprintf("%s%s%s", uqSelectUser, uqWhere, uqFilterPhoneFilter)
	query, args, err := ur.DBList.Backend.Read.In(q, phoneFilter)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.ID == 0 {
		return nil, nil
	}

	return &res, nil
}

func (ur UserDataRepo) IsExistOTP(otp, phone string) (bool, error) {
	var isExist bool

	q := fmt.Sprintf("%s(%s%s%s AND %s)", uqSelectExist, uqSelectUser, uqWhere, uqFilterPhone, uqFilterOTP)

	query, args, err := ur.DBList.Backend.Read.In(q, phone, otp)
	if err != nil {
		return isExist, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}

	return isExist, nil
}

func (ur UserDataRepo) IsExistUser(phone string) (bool, error) {
	var isExist bool

	q := fmt.Sprintf("%s(%s%s%s)", uqSelectExist, uqSelectUser, uqWhere, uqFilterPhoneFilter)

	query, args, err := ur.DBList.Backend.Read.In(q, phone)
	if err != nil {
		return isExist, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}

	return isExist, nil
}

func (ur UserDataRepo) InsertUser(tx *sql.Tx, data du.CreateUser) (int64, error) {
	param := make([]interface{}, 0)

	param = append(param, strings.Title(strings.ToLower(data.Name)))
	param = append(param, data.Status)
	param = append(param, data.Phone)
	param = append(param, data.PhoneFilter)

	var otp *string
	if data.OTP.Valid {
		otp = &data.OTP.String
	}

	param = append(param, otp)

	var otpCreatedAt *time.Time
	if data.OTPCreatedAt.Valid {
		otpCreatedAt = &data.OTPCreatedAt.Time
	}

	param = append(param, otpCreatedAt)
	param = append(param, time.Now().UTC())

	query, args, err := ur.DBList.Backend.Write.In(uqInsertUser, param...)
	if err != nil {
		return 0, err
	}

	query = ur.DBList.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = ur.DBList.Backend.Write.QueryRow(query, args...)
	} else {
		res = tx.QueryRow(query, args...)
	}

	if err != nil {
		return 0, err
	}

	err = res.Err()
	if err != nil {
		return 0, err
	}

	var userID int64
	err = res.Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ur UserDataRepo) VerifyUser(tx *sql.Tx, data du.VerifyUser) error {
	var err error

	q := fmt.Sprintf("%s, %s, %s %s %s AND %s ", uqUpdateUser, uqFilterOTP, uqFilterOTPCreatedAt, uqWhere, uqFilterPhoneFilter, uqFilterOTP)
	query, args, err := ur.DBList.Backend.Read.In(q, cg.UpdatedBySystem, nil, nil, data.PhoneFilter, data.OTP)
	if err != nil {
		return err
	}

	query = ur.DBList.Backend.Write.Rebind(query)
	_, err = ur.DBList.Backend.Write.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserDataRepo) UpdateStatus(tx *sql.Tx, status int, userID int64) error {
	var err error

	q := fmt.Sprintf("%s, %s %s%s", uqUpdateUser, uqFilterStatus, uqWhere, uqFilterUserID)

	query, args, err := ur.DBList.Backend.Read.In(q, cg.UpdatedBySystem, status, userID)
	if err != nil {
		return err
	}

	query = ur.DBList.Backend.Write.Rebind(query)
	_, err = ur.DBList.Backend.Write.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (ur UserDataRepo) UpdateOTP(tx *sql.Tx, otp string, userID int64) error {
	var err error

	q := fmt.Sprintf("%s, %s, %s %s%s", uqUpdateUser, uqFilterOTP, uqFilterOTPCreatedAt, uqWhere, uqFilterUserID)

	query, args, err := ur.DBList.Backend.Read.In(q, cg.UpdatedBySystem, otp, time.Now().UTC(), userID)
	if err != nil {
		return err
	}

	query = ur.DBList.Backend.Write.Rebind(query)
	_, err = ur.DBList.Backend.Write.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
