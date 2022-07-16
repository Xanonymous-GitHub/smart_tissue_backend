package model

type ToiletState string

const (
	SUFFICIENT ToiletState = "sufficient"
	INSUFFICIENT ToiletState = "insufficient"
	DISCONNECTED ToiletState = "disconnected"
	CLEANING ToiletState = "cleaning"
	MAINTAINING ToiletState = "maintaining"
)