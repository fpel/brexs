package domain

type RouteSchema struct {
	Origin  string `json:"origin" valid:"required"`
	Destiny string `json:"destiny" valid:"required"`
	Cost    int    `json:"cost"`
}
