package runexpensecreate

import "your/path/project/domain_myspending/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveExpenseRepo
}
