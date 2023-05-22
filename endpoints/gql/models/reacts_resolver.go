package models

type GQLModel_ReactsResolver struct {
	*Reacts
}

func NewGQLModel_ReactsResolver(r *Reacts) *GQLModel_ReactsResolver {
	return &GQLModel_ReactsResolver{r}
}

func (r *GQLModel_ReactsResolver) LikeCount() float64 {
	return float64(r.Reacts.LikeCount)
}

func (r *GQLModel_ReactsResolver) ConfusedCount() float64 {
	return float64(r.Reacts.ConfusedCount)
}

func (r *GQLModel_ReactsResolver) LoveCount() float64 {
	return float64(r.Reacts.LoveCount)
}

func (r *GQLModel_ReactsResolver) LaughCount() float64 {
	return float64(r.Reacts.LaughCount)
}

func (r *GQLModel_ReactsResolver) SadCount() float64 {
	return float64(r.Reacts.SadCount)
}

func (r *GQLModel_ReactsResolver) CareCount() float64 {
	return float64(r.Reacts.CareCount)
}
