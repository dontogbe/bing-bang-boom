package bingbangboomsvc

type svc struct {
}

func NewSvc() Svc {
	return &svc{}
}
func (s svc) GenerateMapping(id int64) ([]Domain, error) {
	var ds []Domain
	if id%3 == 0 {
		ds = append(ds, Domain_A)
	}
	if id%5 == 0 {
		ds = append(ds, Domain_B)
	}
	return ds, nil
}
