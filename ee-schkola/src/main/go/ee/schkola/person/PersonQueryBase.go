package person

import (
    "context"
    "github.com/eugeis/gee/eh"
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
            ret[i] = e.(*Church)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Church)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Church, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Church)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Church)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Church
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Church
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Church
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Church
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
            ret[i] = e.(*Graduation)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Graduation)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Graduation, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Graduation)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Graduation)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Graduation
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Graduation
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Graduation
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Graduation
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

func (o *QueryRepository) FindByEmail(email string) (ret *Profile, err error) {
    err = eh.QueryNotImplemented("findProfileByEmail")
        
    err = eh.QueryNotImplemented("findProfileByEmail")
    return
}

func (o *QueryRepository) FindByPhone(phone string) (ret *Profile, err error) {
    err = eh.QueryNotImplemented("findProfileByPhone")
        
    err = eh.QueryNotImplemented("findProfileByPhone")
    return
}

func (o *QueryRepository) FindAll() (ret []string, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Profile)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Profile)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Profile, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Profile)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Profile)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Profile
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Profile
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Profile
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Profile
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









