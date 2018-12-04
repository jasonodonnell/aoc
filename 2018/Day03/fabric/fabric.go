package fabric

type Fabric struct {
	Matrix *[1000][1000][]Claim
}

func NewFabric(claims []Claim) *Fabric {
	var fabric [1000][1000][]Claim
	for _, claim := range claims {
		for i := 0; i < claim.Tall; i++ {
			for j := 0; j < claim.Wide; j++ {
				fabric[claim.X+j][claim.Y+i] = append(fabric[claim.X+j][claim.Y+i], claim)
			}
		}
	}
	return &Fabric{
		Matrix: &fabric,
	}
}
