/*
#######
##                  __  _
##       ___  ___  / /_(_)__  ___  ___
##      / _ \/ _ \/ __/ / _ \/ _ \(_-<
##      \___/ .__/\__/_/\___/_//_/___/
##         /_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package options

import (
	"errors"
	"strconv"
	"time"

	"github.com/arnumina/failure"
)

var (
	// ErrBadType AFAIRE.
	ErrBadType = errors.New("bad type")
	// ErrNotFound AFAIRE.
	ErrNotFound = errors.New("not found")
)

// Options AFAIRE.
type Options map[string]interface{}

// New AFAIRE.
func New() Options {
	return make(Options)
}

// Get AFAIRE.
func (o Options) Get(name string) (interface{}, error) {
	if ov, ok := o[name]; ok {
		return ov, nil
	}

	return nil,
		failure.New(ErrNotFound).
			Set("name", name).
			Msg("this option does not exist") //////////////////////////////////////////////////////////////////////////
}

// Boolean AFAIRE.
func (o Options) Boolean(name string) (bool, error) {
	ov, err := o.Get(name)
	if err != nil {
		return false, err
	}

	switch v := ov.(type) {
	case bool:
		return v, nil
	case string:
		return strconv.ParseBool(v)
	default:
		return false,
			failure.New(ErrBadType).
				Set("name", name).
				Msg("this option is not a boolean") ////////////////////////////////////////////////////////////////////
	}
}

// DBoolean AFAIRE.
func (o Options) DBoolean(d bool, name string) (bool, error) {
	v, err := o.Boolean(name)
	if err == nil {
		return v, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return false, err
}

// Integer AFAIRE.
func (o Options) Integer(name string) (int, error) {
	ov, err := o.Get(name)
	if err != nil {
		return 0, err
	}

	switch v := ov.(type) {
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0,
			failure.New(ErrBadType).
				Set("name", name).
				Msg("this option is not an integer") ///////////////////////////////////////////////////////////////////
	}
}

// DInteger AFAIRE.
func (o Options) DInteger(d int, name string) (int, error) {
	v, err := o.Integer(name)
	if err == nil {
		return v, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return 0, err
}

// String AFAIRE.
func (o Options) String(name string) (string, error) {
	ov, err := o.Get(name)
	if err != nil {
		return "", err
	}

	if v, ok := ov.(string); ok {
		return v, nil
	}

	return "",
		failure.New(ErrBadType).
			Set("name", name).
			Msg("this option is not a string") /////////////////////////////////////////////////////////////////////////
}

// DString AFAIRE.
func (o Options) DString(d, name string) (string, error) {
	v, err := o.String(name)
	if err == nil {
		return v, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return "", err
}

// Duration AFAIRE.
func (o Options) Duration(name string) (time.Duration, error) {
	ov, err := o.Get(name)
	if err != nil {
		return 0, err
	}

	s, ok := ov.(string)
	if ok {
		v, err := time.ParseDuration(s)
		if err == nil {
			return v, nil
		}
	}

	return 0,
		failure.New(ErrBadType).
			Set("name", name).
			Msg("this option is not a duration") ///////////////////////////////////////////////////////////////////////
}

// DDuration AFAIRE.
func (o Options) DDuration(d time.Duration, name string) (time.Duration, error) {
	v, err := o.Duration(name)
	if err == nil {
		return v, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return 0, err
}

/*
######################################################################################################## @(°_°)@ #######
*/
