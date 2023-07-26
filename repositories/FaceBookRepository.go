package repositories

type FaceBookRepository interface{}

type facebookRepository struct{}

func NewFaceBookRepository() FaceBookRepository {
	return &facebookRepository{}
}
