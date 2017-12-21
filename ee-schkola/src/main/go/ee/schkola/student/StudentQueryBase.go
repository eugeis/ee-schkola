package student

import (
    "context"
    "github.com/looplab/eventhorizon"
)
type QueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewAttendanceQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []*Attendance, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Attendance, len(result))
		for i, e := range result {
            ret[i] = e.(*Attendance)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Attendance, len(result))
		for i, e := range result {
            ret[i] = e.(*Attendance)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Attendance, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Attendance)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Attendance)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Attendance
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Attendance
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Attendance
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Attendance
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

func NewCourseQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []*Course, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Course, len(result))
		for i, e := range result {
            ret[i] = e.(*Course)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Course, len(result))
		for i, e := range result {
            ret[i] = e.(*Course)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Course, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Course)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Course)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Course
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Course
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Course
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Course
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

func NewGradeQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []*Grade, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Grade, len(result))
		for i, e := range result {
            ret[i] = e.(*Grade)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Grade, len(result))
		for i, e := range result {
            ret[i] = e.(*Grade)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Grade, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Grade)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Grade)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Grade
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Grade
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Grade
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Grade
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

func NewGroupQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []*Group, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Group, len(result))
		for i, e := range result {
            ret[i] = e.(*Group)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Group, len(result))
		for i, e := range result {
            ret[i] = e.(*Group)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *Group, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Group)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Group)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*Group
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*Group
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Group
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *Group
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

func NewSchoolApplicationQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []*SchoolApplication, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*SchoolApplication, len(result))
		for i, e := range result {
            ret[i] = e.(*SchoolApplication)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*SchoolApplication, len(result))
		for i, e := range result {
            ret[i] = e.(*SchoolApplication)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *SchoolApplication, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*SchoolApplication)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*SchoolApplication)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*SchoolApplication
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*SchoolApplication
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *SchoolApplication
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *SchoolApplication
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

func NewSchoolYearQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *QueryRepository) {
    ret = &QueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *QueryRepository) FindAll() (ret []*SchoolYear, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*SchoolYear, len(result))
		for i, e := range result {
            ret[i] = e.(*SchoolYear)
		}
    }
        
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*SchoolYear, len(result))
		for i, e := range result {
            ret[i] = e.(*SchoolYear)
		}
    }
    return
}

func (o *QueryRepository) FindById(id eventhorizon.UUID) (ret *SchoolYear, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*SchoolYear)
    }
        
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*SchoolYear)
    }
    return
}

func (o *QueryRepository) CountAll() (ret int, err error) {
    var result []*SchoolYear
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
        
    var result []*SchoolYear
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *QueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *SchoolYear
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
        
    var result *SchoolYear
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









