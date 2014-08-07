package models

// A GitHub repository
type Repository struct {
	Id        int64
	Name      string
	Url       string
	Languages Languages
}

func (r *Repository) LanguageDistribution() LanguageDistribution {
	total := int64(0)
	distribution := LanguageDistribution{}
	for _, value := range r.Languages {
		total = total + value
	}

	for key, value := range r.Languages {
		distribution[key] = float64(value) * 100.0 / float64(total)
	}

	return distribution
}
