package repositories

type AdminRepository interface {
}

type AdminRepositoryImpl struct {
}

func NewAdminRepositoryImpl() AdminRepository {
	return &AdminRepositoryImpl{}
}
