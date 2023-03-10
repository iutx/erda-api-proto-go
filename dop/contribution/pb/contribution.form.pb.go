// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: contribution.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*PersonalContribution)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Contribution)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Indicator)(nil)
var _ urlenc.URLValuesUnmarshaler = (*IndicatorData)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UserRank)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetPersonalContributionRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetPersonalContributionResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetActiveRankRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetActiveRankRequestResponse)(nil)

// PersonalContribution implement urlenc.URLValuesUnmarshaler.
func (m *PersonalContribution) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Contribution{}
				}
			case "data.events":
				if m.Data == nil {
					m.Data = &Contribution{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Events = val
			case "data.commits":
				if m.Data == nil {
					m.Data = &Contribution{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Commits = val
			case "data.cases":
				if m.Data == nil {
					m.Data = &Contribution{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Cases = val
			case "indicators":
				if m.Indicators == nil {
					m.Indicators = &Indicator{}
				}
			case "indicators.max":
				if m.Indicators == nil {
					m.Indicators = &Indicator{}
				}
				list := make([]uint64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseUint(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Indicators.Max = list
			case "indicators.title":
				if m.Indicators == nil {
					m.Indicators = &Indicator{}
				}
				m.Indicators.Title = vals
			}
		}
	}
	return nil
}

// Contribution implement urlenc.URLValuesUnmarshaler.
func (m *Contribution) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "events":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Events = val
			case "commits":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Commits = val
			case "cases":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Cases = val
			}
		}
	}
	return nil
}

// Indicator implement urlenc.URLValuesUnmarshaler.
func (m *Indicator) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "max":
				list := make([]uint64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseUint(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Max = list
			case "title":
				m.Title = vals
			}
		}
	}
	return nil
}

// IndicatorData implement urlenc.URLValuesUnmarshaler.
func (m *IndicatorData) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				list := make([]uint64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseUint(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Data = list
			}
		}
	}
	return nil
}

// UserRank implement urlenc.URLValuesUnmarshaler.
func (m *UserRank) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "rank":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Rank = val
			case "value":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Value = val
			}
		}
	}
	return nil
}

// GetPersonalContributionRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetPersonalContributionRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgID":
				m.OrgID = vals[0]
			case "userID":
				m.UserID = vals[0]
			}
		}
	}
	return nil
}

// GetPersonalContributionResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetPersonalContributionResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &PersonalContribution{}
				}
			case "data.data":
				if m.Data == nil {
					m.Data = &PersonalContribution{}
				}
				if m.Data.Data == nil {
					m.Data.Data = &Contribution{}
				}
			case "data.data.events":
				if m.Data == nil {
					m.Data = &PersonalContribution{}
				}
				if m.Data.Data == nil {
					m.Data.Data = &Contribution{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Data.Events = val
			case "data.data.commits":
				if m.Data == nil {
					m.Data = &PersonalContribution{}
				}
				if m.Data.Data == nil {
					m.Data.Data = &Contribution{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Data.Commits = val
			case "data.data.cases":
				if m.Data == nil {
					m.Data = &PersonalContribution{}
				}
				if m.Data.Data == nil {
					m.Data.Data = &Contribution{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Data.Cases = val
			case "data.indicators":
				if m.Data == nil {
					m.Data = &PersonalContribution{}
				}
				if m.Data.Indicators == nil {
					m.Data.Indicators = &Indicator{}
				}
			case "data.indicators.max":
				if m.Data == nil {
					m.Data = &PersonalContribution{}
				}
				if m.Data.Indicators == nil {
					m.Data.Indicators = &Indicator{}
				}
				list := make([]uint64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseUint(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Data.Indicators.Max = list
			case "data.indicators.title":
				if m.Data == nil {
					m.Data = &PersonalContribution{}
				}
				if m.Data.Indicators == nil {
					m.Data.Indicators = &Indicator{}
				}
				m.Data.Indicators.Title = vals
			}
		}
	}
	return nil
}

// GetActiveRankRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetActiveRankRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgID":
				m.OrgID = vals[0]
			}
		}
	}
	return nil
}

// GetActiveRankRequestResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetActiveRankRequestResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "userIDs":
				m.UserIDs = vals
			}
		}
	}
	return nil
}
