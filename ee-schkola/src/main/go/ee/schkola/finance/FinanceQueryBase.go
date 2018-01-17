package finance

import (
    "context"
    "github.com/looplab/eventhorizon"
)
type QueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []string, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Expense)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Expense)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Expense, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Expense)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Expense)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Expense
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Expense
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Expense
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Expense
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *QueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
        
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *QueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
        
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type QueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []string, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*ExpensePurpose)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*ExpensePurpose)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *ExpensePurpose, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*ExpensePurpose)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*ExpensePurpose)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*ExpensePurpose
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*ExpensePurpose
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *ExpensePurpose
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *ExpensePurpose
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *QueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
        
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *QueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
        
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type QueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []string, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Fee)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Fee)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Fee, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Fee)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Fee)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Fee
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Fee
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Fee
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Fee
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *QueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
        
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *QueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
        
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type QueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []string, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*FeeKind)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*FeeKind)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *FeeKind, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*FeeKind)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*FeeKind)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*FeeKind
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*FeeKind
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *FeeKind
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *FeeKind
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *QueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
        
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *QueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
        
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}









