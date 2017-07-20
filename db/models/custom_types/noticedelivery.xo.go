package custom_types

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// NoticeDelivery is the 'notice_delivery' enum type from schema 'public'.
type NoticeDelivery uint16

const (
	// NoticeDeliveryNone is the 'none' NoticeDelivery.
	NoticeDeliveryNone = NoticeDelivery(1)

	// NoticeDeliverySite is the 'site' NoticeDelivery.
	NoticeDeliverySite = NoticeDelivery(2)

	// NoticeDeliveryWeekly is the 'weekly' NoticeDelivery.
	NoticeDeliveryWeekly = NoticeDelivery(3)

	// NoticeDeliveryDaily is the 'daily' NoticeDelivery.
	NoticeDeliveryDaily = NoticeDelivery(4)

	// NoticeDeliveryAsync is the 'async' NoticeDelivery.
	NoticeDeliveryAsync = NoticeDelivery(5)
)

// String returns the string value of the NoticeDelivery.
func (nd NoticeDelivery) String() string {
	var enumVal string

	switch nd {
	case NoticeDeliveryNone:
		enumVal = "none"

	case NoticeDeliverySite:
		enumVal = "site"

	case NoticeDeliveryWeekly:
		enumVal = "weekly"

	case NoticeDeliveryDaily:
		enumVal = "daily"

	case NoticeDeliveryAsync:
		enumVal = "async"
	}

	return enumVal
}

// MarshalText marshals NoticeDelivery into text.
func (nd NoticeDelivery) MarshalText() ([]byte, error) {
	return []byte(nd.String()), nil
}

// UnmarshalText unmarshals NoticeDelivery from text.
func (nd *NoticeDelivery) UnmarshalText(text []byte) error {
	switch string(text) {
	case "none":
		*nd = NoticeDeliveryNone

	case "site":
		*nd = NoticeDeliverySite

	case "weekly":
		*nd = NoticeDeliveryWeekly

	case "daily":
		*nd = NoticeDeliveryDaily

	case "async":
		*nd = NoticeDeliveryAsync

	default:
		return errors.New("invalid NoticeDelivery")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for NoticeDelivery.
func (nd NoticeDelivery) Value() (driver.Value, error) {
	return nd.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for NoticeDelivery.
func (nd *NoticeDelivery) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid NoticeDelivery")
	}

	return nd.UnmarshalText(buf)
}

type NullNoticeDelivery struct {
	NoticeDelivery NoticeDelivery
	Valid          bool
}

func (nv *NullNoticeDelivery) Scan(value interface{}) error {
	if value == nil {
		nv.NoticeDelivery, nv.Valid = NoticeDelivery(0), false
		return nil
	}
	err := nv.NoticeDelivery.Scan(value)
	if err != nil {
		nv.Valid = false
		return err
	} else {
		nv.Valid = true
		return nil
	}
}

func (nv NullNoticeDelivery) Value() (driver.Value, error) {
	if !nv.Valid {
		return nil, nil
	}
	return nv.NoticeDelivery.Value()
}

func NoticeDeliveryRandom() NoticeDelivery {
	return NoticeDeliveryAsync

}

func NullNoticeDeliveryRandom() NullNoticeDelivery {
	return NullNoticeDelivery{NoticeDeliveryAsync, true}

}
