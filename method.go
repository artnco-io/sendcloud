package sendcloud

import "strconv"

type Method struct {
	ID          int64
	Name        string
	CarrierCode string
	Amount      int64
	MinWeight   int64
	MaxWeight   int64
	Countries   []string
}

type MethodListResponseContainer struct {
	ShippingMethods []MethodResponse `json:"shipping_methods"`
}

type MethodResponseContainer struct {
	ShippingMethod MethodResponse `json:"shipping_method"`
}

type MethodResponse struct {
	ServicePointInput string            `json:"service_point_input"`
	MaxWeight         string            `json:"max_weight"`
	Name              string            `json:"name"`
	Carrier           string            `json:"carrier"`
	Countries         []CountryResponse `json:"countries"`
	MinWeight         string            `json:"min_weight"`
	ID                int64             `json:"id"`
	Price             float64           `json:"price"`
}

type CountryResponse struct {
	Iso2  string  `json:"iso_2"`
	Iso3  string  `json:"iso_3"`
	ID    int     `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}

func (a *MethodListResponseContainer) GetResponse() interface{} {
	var methods []*Method
	for _, sm := range a.ShippingMethods {
		method := sm.ToMethod()
		methods = append(methods, method)

	}
	return methods
}

func (m *MethodResponseContainer) GetResponse() interface{} {
	method := m.ShippingMethod.ToMethod()
	return method
}

func (sm *MethodResponse) ToMethod() *Method {
	maxWeightFloat, _ := strconv.ParseFloat(sm.MaxWeight, 64)
	maxWeight := int64(maxWeightFloat * 1000)
	minWeightFloat, _ := strconv.ParseFloat(sm.MinWeight, 64)
	minWeight := int64(minWeightFloat * 1000)

	method := &Method{
		ID:          sm.ID,
		Name:        sm.Name,
		CarrierCode: sm.Carrier,
		Amount:      int64(sm.Price) * 100,
		MinWeight:   minWeight,
		MaxWeight:   maxWeight,
	}
	for _, c := range sm.Countries {
		method.Countries = append(method.Countries, c.Iso2)
	}

	return method
}