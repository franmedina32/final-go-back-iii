package domain

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Paciente struct {
	ID        int        `json:"id"`
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Domicilio string     `json:"domicilio"`
	Dni       string     `json:"dni"`
	FechaAlta CustomTime `json:"fecha_alta"`
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	str := string(b)
	str = strings.Trim(str, "\"")
	if str == "null" {
		ct.Time = time.Time{}
		return nil
	}
	layout := "2006-01-02 15:04:05"
	ct.Time, err = time.Parse(layout, str)
	return err
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	formattedTime := ct.Time.Format("2006-01-02 15:04:05")
	return []byte(`"` + formattedTime + `"`), nil
}

func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		ct.Time = v
	case []byte:
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return fmt.Errorf("error parsing time: %v", err)
		}
		ct.Time = parsedTime
	default:
		return fmt.Errorf("unsupported type for CustomTime: %T", value)
	}
	return nil
}
