package repo

import (
	"context"
	. "taylor-ai-server/internal/domain"
	"taylor-ai-server/internal/pkg/connection"

	"github.com/doug-martin/goqu/v9"
)

const userTable = "users"

type UserRepo struct {
	*operators
}

func NewUserRepo() *UserRepo {
	db := connection.DB()
	return &UserRepo{
		operators: newOperators(db, userTable),
	}
}

func (r *UserRepo) find(ctx context.Context, wheres ...goqu.Expression) (User, error) {
	var e userEntity
	found, err := r.selector().Where(wheres...).ScanStructContext(ctx, &e)
	if err != nil {
		return User{}, err
	}
	if !found {
		return User{}, ErrNotFound
	}
	return r.toDomain(e), nil
}

func (r *UserRepo) Find(ctx context.Context, id string) (User, error) {
	return r.find(ctx, goqu.C("id").Eq(id))
}

func (r *UserRepo) FindByUsername(ctx context.Context, username string) (User, error) {
	return r.find(ctx, goqu.C("username").Eq(username))
}

func (r *UserRepo) Create(ctx context.Context, m User) error {
	e := r.toEntity(m)
	_, err := r.inserter().Rows(e).Executor().ExecContext(ctx)
	return err
}

func (r *UserRepo) Update(ctx context.Context, m User) error {
	e := r.toEntity(m)
	_, err := r.updater().Set(e).Where(goqu.C("id").Eq(e.ID)).Executor().ExecContext(ctx)
	return err
}

func (r *UserRepo) toDomain(e userEntity) User {
	return User{
		ID:             e.ID,
		Username:       e.Username,
		Nickname:       e.Nickname,
		DeviceID:       e.DeviceID,
		Password:       e.Password,
		Email:          e.Email,
		Avatar:         e.Avatar,
		AppToken:       e.AppToken,
		WebToken:       e.WebToken,
		RegisterTime:   e.RegisterTime,
		RegisterIP:     e.RegisterIP,
		Lang:           e.Lang,
		LoginTime:      e.LoginTime,
		IsPro:          intToBool(e.IsPro),
		ProType:        e.ProType,
		ProIsAutoRenew: intToBool(e.ProIsAutoRenew),
		ProStart:       e.ProStart,
		ProEnd:         e.ProEnd,
		CreateTime:     e.CreateTime,
		UpdateTime:     e.UpdateTime,
		DeleteTime:     e.DeleteTime,
	}
}

func (r *UserRepo) toEntity(d User) userEntity {
	return userEntity{
		ID:             d.ID,
		Username:       d.Username,
		Nickname:       d.Nickname,
		DeviceID:       d.DeviceID,
		Password:       d.Password,
		Email:          d.Email,
		Avatar:         d.Avatar,
		AppToken:       d.AppToken,
		WebToken:       d.WebToken,
		RegisterTime:   d.RegisterTime,
		RegisterIP:     d.RegisterIP,
		Lang:           d.Lang,
		LoginTime:      d.LoginTime,
		IsPro:          boolToInt(d.IsPro),
		ProType:        d.ProType,
		ProIsAutoRenew: boolToInt(d.ProIsAutoRenew),
		ProStart:       d.ProStart,
		ProEnd:         d.ProEnd,
		CreateTime:     d.CreateTime,
		UpdateTime:     d.UpdateTime,
		DeleteTime:     d.DeleteTime,
	}
}

func (r *UserRepo) toDomains(es userEntities) Users {
	ms := make([]User, 0, len(es))
	for _, e := range es {
		ms = append(ms, r.toDomain(e))
	}
	return ms
}

func (r *UserRepo) toEntities(ds Users) userEntities {
	es := make(userEntities, 0, len(ds))
	for _, d := range ds {
		es = append(es, r.toEntity(d))
	}
	return es
}

type userEntity struct {
	ID             string `db:"id"`
	Username       string `db:"username"`
	Nickname       string `db:"nickname"`
	DeviceID       string `db:"device_id"`
	Password       string `db:"password"`
	Email          string `db:"email"`
	Avatar         string `db:"avatar"`
	AppToken       string `db:"app_token"`
	WebToken       string `db:"web_token"`
	RegisterTime   int64  `db:"register_time"`
	RegisterIP     string `db:"register_ip"`
	Lang           string `db:"lang"`
	LoginTime      int64  `db:"login_time"`
	IsPro          int    `db:"is_pro"`
	ProType        int    `db:"pro_type"`
	ProIsAutoRenew int    `db:"pro_is_auto_renew"`
	ProStart       int64  `db:"pro_start"`
	ProEnd         int64  `db:"pro_end"`
	CreateTime     int64  `db:"create_time"`
	UpdateTime     int64  `db:"update_time"`
	DeleteTime     int64  `db:"delete_time"`
}

type userEntities []userEntity
