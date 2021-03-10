package models

type Model interface {
	Get() error
	Add() (string, error)
	Update() error
	Delete() error
}
