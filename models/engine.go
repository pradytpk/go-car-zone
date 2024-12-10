package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID       uuid.UUID `json:"engine_id"`
	Displacement   int64     `json:"displacement"`
	NoOfCyclinders int64     `json:"no_of_cyclinders"`
	CarRange       int64     `json:"car_range"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type EngineRequest struct {
	Displacement   int64 `json:"displacement"`
	NoOfCyclinders int64 `json:"no_of_cyclinders"`
	CarRange       int64 `json:"car_range"`
}

func ValidateEngineRequest(engineRequest EngineRequest) error {
	if err := validateDisplacement(engineRequest.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCyclinders(engineRequest.NoOfCyclinders); err != nil {
		return err
	}
	if err := validateCarRange(engineRequest.CarRange); err != nil {
		return err
	}
	return nil
}

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacment must be greater than zero")
	}
	return nil
}

func validateNoOfCyclinders(noOfCyclinders int64) error {
	if noOfCyclinders <= 0 {
		return errors.New("no of cylinders must be greater than zero")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("car range must be greater than zero")
	}
	return nil
}
