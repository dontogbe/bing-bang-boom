package bingbangboomsvc

type Domain string

const (
	Domain_A Domain = "A"
	Domain_B Domain = "B"
)

type Svc interface {
	GenerateMapping(id int64) ([]Domain,error)
}

