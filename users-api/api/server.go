package api

import (
	"users-api/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserServer struct {
	UIDGenerator *utils.UniqueIDGenerator
	Argon2Params *utils.Argon2IDParams
	DatabasePool *pgxpool.Pool
}
