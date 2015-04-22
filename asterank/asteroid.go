// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package asterank

// Asteroid defines a single asteroid.
//
// Note: The meaning of all these fields is not entirely clear.
// More information may be available in the JPL Small-Body-Database at
// http://ssd.jpl.nasa.gov/sbdb.cgi
type Asteroid map[string]interface{}

/*
This is not a practical way to implement an Asteroid.

Many of the fields have inconsistent data types in the response data.
Which make sus end up with a lot of `interface{}` typed fields. Might as well
just go with a map.

type Asteroid struct {
	Price            float64 `json:"price"`
	Profit           float64 `json:"profit"`
	DeltaV           float64 `json:"dv"`
	Aphelion         float64 `json:"a"`
	Eccentricity     float64 `json:"e"`
	Inclination      float64 `json:"i"`
	H                float64 `json:"H"`
	N                float64 `json:"n"`
	Q                float64 `json:"q"`
	W                float64 `json:"w"`
	Epoch            float64 `json:"epoch"`
	EpochCal         float64 `json:"epoch_cal"`
	EpochMJD         float64 `json:"epoch_mjd"`
	Ad               float64 `json:"ad"`
	Rms              float64 `json:"rms"`
	Closeness        float64 `json:"closeness"`
	Saved            float64 `json:"saved"`
	MoidLD           float64 `json:"moid_ld"`
	EstDiameter      float64 `json:"est_diameter"`
	Per              float64 `json:"per"`
	DataArc          float64 `json:"data_arc"`
	Score            float64 `json:"score"`
	PerY             float64 `json:"per_y"`
	ConditionCode    float64 `json:"condition_code"`
	ObservationsUsed float64 `json:"n_obs_used"`
	Moid             float64 `json:"moid"`
	TpCal            float64 `json:"tp_cal"`
	TJup             float64 `json:"t_jup"`
	Om               float64 `json:"om"`
	Ma               float64 `json:"ma"`
	Tp               float64 `json:"tp"`
	SpkId            float64 `json:"spkid"`
	SigmaPer         float64 `json:"sigma_per"`
	SigmaAD          float64 `json:"sigma_ad"`
	SigmaW           float64 `json:"sigma_w"`
	SigmaI           float64 `json:"sigma_i"`
	SigmaTP          float64 `json:"sigma_tp"`
	SigmaQ           float64 `json:"sigma_q"`
	SigmaN           float64 `json:"sigma_n"`
	SigmaA           float64 `json:"sigma_a"`
	SigmaOm          float64 `json:"sigma_om"`
	SigmaE           float64 `json:"sigma_e"`
	SigmaMa          float64 `json:"sigma_ma"`

	Id                     string `json:"id"`
	Name                   string `json:"name"`
	FullName               string `json:"full_name"`
	Class                  string `json:"class"`
	Producer               string `json:"producer"`
	HSigma                 string `json:"H_sigma"`
	G                      string `json:"G"`
	DT                     string `json:"DT"`
	PC                     string `json:"PC"`
	IR                     string `json:"IR"`
	A1                     string `json:"A1"`
	A2                     string `json:"A2"`
	A3                     string `json:"A3"`
	K1                     string `json:"K1"`
	K2                     string `json:"K2"`
	M1                     string `json:"M1"`
	M2                     string `json:"M2"`
	TwoBody                string `json:"two_body"`
	Equinox                string `json:"equinox"`
	Pha                    string `json:"pha"`
	Neo                    string `json:"neo"`
	OrbitId                string `json:"orbit_id"`
	ProvDes                string `json:"prov_des"`
	Spec                   string `json:"spec"`
	SpecT                  string `json:"spec_T"`
	SpecB                  string `json:"spec_B"`
	Extent                 string `json:"extent"`
	Prefix                 string `json:"prefix"`
	FirstObservation       string `json:"first_obs"`
	LastObservation        string `json:"last_obs"`
	NumDelObservationsUsed string `json:"n_del_obs_used"`
	NumDopObservationsUsed string `json:"n_dop_obs_used"`

	// Fields with varying data types --sometimes string, sometimes float64
	Pdes          interface{} `json:"pdes"`
	RotPer        interface{} `json:"rot_per"`
	Diameter      interface{} `json:"diameter"`
	DiameterSigma interface{} `json:"diameter_sigma"`
	Albedo        interface{} `json:"albedo"`
	GM            interface{} `json:"GM"`
	UB            interface{} `json:"UB"`
	BV            interface{} `json:"BV"`
}
*/
