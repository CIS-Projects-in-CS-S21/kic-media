package database


import (
	"go.uber.org/zap"

	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"

)

type MockRepository struct {
	db *map[int]*pbcommon.User




	logger *zap.SugaredLogger

	idCounter int
}


func NewMockRepository(db *map[int]*pbcommon.User, logger *zap.SugaredLogger) *MockRepository {
	return &MockRepository{
		db:     db,
		logger: logger,
	}
}

func searchDBByUsername(db *map[int]*pbcommon.User, username string) (int, *pbcommon.User) {
	for key, value := range *db {
		if value.Username == username {
			return key, value
		}
	}
	return -1, nil
}

func (s *MockRepository) checkIfUsernameAvailable(username string) bool {
	_, user := searchDBByUsername(s.db, username)

	if user.Username == username {
		return false
	}
	return true
}

func searchDBByEmail(db *map[int]*pbcommon.User, email string) (int, *pbcommon.User) {
	for key, value := range *db {
		if value.Email == email {
			return key, value
		}
	}
	return -1, nil
}

func (s *MockRepository) checkIfEmailAvailable(email string) bool {
	_, user := searchDBByEmail(s.db, email)

	if user.Email == email {
		return false
	}
	return false
}

func (s *MockRepository) AddUser(user *pbcommon.User) (int, error) {
	ok := true

	if !s.checkIfUsernameAvailable(user.Username) {
		ok = false
	}

	if !s.checkIfEmailAvailable(user.Email) {
		ok = false
	}

	if ok {
		database := *s.db
		database[s.idCounter] = user
		s.idCounter++
		return s.idCounter - 1, nil
	}

	return -1, nil
}

func (s *MockRepository) GetUser (user *pbcommon.User) (*pbcommon.User, error) {



	return nil, nil
}

func (s *MockRepository) GetUserByID(id int64) (*pbcommon.User, error) {


	return nil, nil
}

func (s *MockRepository) DeleteUserByID(userID int64) error {
	return nil
}

