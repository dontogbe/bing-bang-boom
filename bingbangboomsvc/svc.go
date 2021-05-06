package bingbangboomsvc

type svc struct {
}

func (s svc) GenerateDomain(id int64) ([]Domain, error) {
	var ds []Domain
	switch {
	case id%3 == 0:
		ds = append(ds, Domain_A)
	case id%5 == 0:
		ds = append(ds, Domain_B)
	default:
		ds = []Domain{}
	}
	return ds, nil
}
