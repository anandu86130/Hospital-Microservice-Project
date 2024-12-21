package interfaces

import (
	"github.com/anandu86130/Hospital-admin-service/internal/model"
)

type AdminRepoInter interface {
	FindAdminByEmail(email string) (*model.Admin, error)
}
 