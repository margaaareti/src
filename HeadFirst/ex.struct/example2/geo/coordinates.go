package geo

import "errors"

type Coordinates struct {
	latitude   float64
	longtitude float64
}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

func (c *Coordinates) Longtitude() float64 {
	return c.longtitude
}

func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}

func (c *Coordinates) SetLongtitude(longtitude float64) error {
	if longtitude < -180 || longtitude > 180 {
		return errors.New("invalid longtitude")
	}
	c.longtitude = longtitude
	return nil
}
