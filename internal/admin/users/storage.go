package users

import "context"

type Repository interface {
	GetAllData(ctx context.Context, req PaginationDTO) (DataDTO, error)
	AddData(ctx context.Context, dto ReqUser) (int, error)
	UpdateData(ctx context.Context, dto ReqUser) error
	DeleteData(ctx context.Context, id int) error

	UpdateAgeData(ctx context.Context, dto ResUserAge) error
	UpdateGenderData(ctx context.Context, dto ResUserGender) error
	UpdateNationalizeData(ctx context.Context, dto ResUserCountry) error
}
