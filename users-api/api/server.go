package api

import "users-api/utils"

type UserServer struct {
	UIDGenerator *utils.UniqueIDGenerator
	Argon2Params *utils.Argon2IDParams
}
