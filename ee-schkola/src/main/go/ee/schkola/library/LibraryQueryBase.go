package library

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

func (o *QueryRepository) FindByTitle(title string) (ret *Book, err error) {
    err = eh.QueryNotImplemented("findBookByTitle")
        
    err = eh.QueryNotImplemented("findBookByTitle")
    return
}

func (o *QueryRepository) FindByPattern(pattern string) (ret *Book, err error) {
    err = eh.QueryNotImplemented("findBookByPattern")
        
    err = eh.QueryNotImplemented("findBookByPattern")
    return
}

func (o *QueryRepository) FindAll() (ret []string, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Book)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]string, len(result))
		for i, e := range result {
            ret[i] = e.(*Book)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Book, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Book)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Book)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Book
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Book
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Book
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Book
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









