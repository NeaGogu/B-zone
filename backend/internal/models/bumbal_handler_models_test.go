package models

import (
	"testing"
)

func TestActivityModelBumbal_SetId(t *testing.T) {
	t.Run("Cycle of Activity Models", func(t *testing.T) {

		//Test without Nil values
		activityModel := NewActivityModel()
		addressAppliedModel := NewAddressAppliedModel()
		addressModel := NewAddressModel()

		addressAppliedModel.SetLongitude("1")
		addressAppliedModel.SetZipcode("1")
		addressAppliedModel.SetLatitude("1")

		if addressAppliedModel.GetLongitude() != "1" {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModel.GetLongitude(), 1)
		}
		if addressAppliedModel.GetLatitude() != "1" {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModel.GetLatitude(), 1)
		}
		if addressAppliedModel.GetZipcode() != "1" {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModel.GetZipcode(), 1)
		}

		addressModel.SetLongitude("1")
		addressModel.SetZipcode("1")
		addressModel.SetLatitude("1")

		if addressModel.GetLongitude() != "1" {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressModel.GetLongitude(), 1)
		}
		if addressModel.GetLatitude() != "1" {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressModel.GetLatitude(), 1)
		}
		if addressModel.GetZipcode() != "1" {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressModel.GetZipcode(), 1)
		}

		activityModel.SetId("1")
		if activityModel.GetId() != "1" {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", activityModel.GetId(), 1)
		}
		activityModel.SetAddressApplied(*addressAppliedModel)
		activityModel.SetDepotAddress(*addressModel)

		//Test objects with NIL values
		activityModelNil := NewActivityModel()
		addressAppliedModelNil := NewAddressAppliedModel()
		addressModelNil := NewAddressModel()

		var ret string
		if addressAppliedModelNil.GetLongitude() != ret {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModelNil.GetLongitude(), ret)
		}
		if addressAppliedModelNil.GetLatitude() != ret {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModelNil.GetLatitude(), ret)
		}
		if addressAppliedModelNil.GetZipcode() != ret {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModelNil.GetZipcode(), ret)
		}

		if addressModelNil.GetLongitude() != ret {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModelNil.GetLongitude(), ret)
		}
		if addressModelNil.GetLatitude() != ret {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModelNil.GetLatitude(), ret)
		}
		if addressModelNil.GetZipcode() != ret {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", addressAppliedModelNil.GetZipcode(), ret)
		}

		if activityModelNil.GetId() != ret {
			t.Errorf("ActivityListResponseBumbal CountUnfiltered = %v, want %v", activityModel.GetId(), ret)
		}

	})
}
