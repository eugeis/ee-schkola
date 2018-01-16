package shared






type PersonName struct {
    First string `json:"first" eh:"optional"`
    Last string `json:"last" eh:"optional"`
}

func New@@EMPTY@@() (ret *PersonName) {
    ret = &PersonName{}
    return
}





