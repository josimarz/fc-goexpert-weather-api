package location

var (
	Regions = map[string]string{
		"AC": "Acre",
		"AL": "Alagoas",
		"AP": "Amapá",
		"AM": "Amazonas",
		"BA": "Bahia",
		"CE": "Ceará",
		"ES": "Espírito Santo",
		"GO": "Goiás",
		"MA": "Maranhão",
		"MT": "Mato Grosso",
		"MS": "Mato Grosso do Sul",
		"MG": "Minas Gerais",
		"PA": "Pará",
		"PB": "Paraíba",
		"PR": "Paraná",
		"PE": "Pernambuco",
		"PI": "Piauí",
		"RJ": "Rio de Janeiro",
		"RN": "Rio Grande do Norte",
		"RS": "Rio Grande do Sul",
		"RO": "Rondônia",
		"RR": "Roraima",
		"SC": "Santa Catarina",
		"SP": "São Paulo",
		"SE": "Sergipe",
		"TO": "Tocantins",
		"DF": "Distrito Federal",
	}
)

type Location struct {
	City   string
	Region string
}

func (l *Location) Equals(other *Location) bool {
	return l.Region == other.Region && l.City == other.City
}
