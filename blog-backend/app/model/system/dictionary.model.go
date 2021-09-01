package system

import "blog-backend/app/model"

type DictionaryHasManyDetails struct {
	*model.Dictionaries `json:"dictionary"`
	DictionaryDetails   []*model.DictionaryDetails `json:"details"`
}
