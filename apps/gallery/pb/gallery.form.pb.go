// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: gallery.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*ListOpusReq)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListOpusResp)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListOpusRespData)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListOpusRespDataItem)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Readme)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Installation)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Form)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Line)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListOpusTypesRespData)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CatalogInfo)(nil)
var _ urlenc.URLValuesUnmarshaler = (*OpusType)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PutOnArtifactsReq)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ArtifactsInstallation)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PutOffArtifactsReq)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PutOnExtensionsReq)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListOpusVersionsReq)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListOpusVersionsResp)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListOpusVersionsRespData)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListOpusVersionRespDataVersion)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PutOnOpusResp)(nil)
var _ urlenc.URLValuesUnmarshaler = (*I18N)(nil)

// ListOpusReq implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusReq) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "type":
				m.Type = vals[0]
			case "name":
				m.Name = vals[0]
			case "keyword":
				m.Keyword = vals[0]
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.PageNo = int32(val)
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.PageSize = int32(val)
			}
		}
	}
	return nil
}

// ListOpusResp implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusResp) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &ListOpusRespData{}
				}
			case "data.total":
				if m.Data == nil {
					m.Data = &ListOpusRespData{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.Total = int32(val)
			case "userIDs":
				m.UserIDs = vals
			}
		}
	}
	return nil
}

// ListOpusRespData implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusRespData) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "total":
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Total = int32(val)
			}
		}
	}
	return nil
}

// ListOpusRespDataItem implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusRespDataItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "createdAt":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
			case "createdAt.seconds":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreatedAt.Seconds = val
			case "createdAt.nanos":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.CreatedAt.Nanos = int32(val)
			case "updatedAt":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "updatedAt.seconds":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdatedAt.Seconds = val
			case "updatedAt.nanos":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.UpdatedAt.Nanos = int32(val)
			case "orgID":
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.OrgID = uint32(val)
			case "orgName":
				m.OrgName = vals[0]
			case "creatorID":
				m.CreatorID = vals[0]
			case "updaterID":
				m.UpdaterID = vals[0]
			case "type":
				m.Type = vals[0]
			case "typeName":
				m.TypeName = vals[0]
			case "name":
				m.Name = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			case "summary":
				m.Summary = vals[0]
			case "catalog":
				m.Catalog = vals[0]
			case "catalogName":
				m.CatalogName = vals[0]
			case "logoURL":
				m.LogoURL = vals[0]
			}
		}
	}
	return nil
}

// Readme implement urlenc.URLValuesUnmarshaler.
func (m *Readme) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "lang":
				m.Lang = vals[0]
			case "langName":
				m.LangName = vals[0]
			case "source":
				m.Source = vals[0]
			case "text":
				m.Text = vals[0]
			}
		}
	}
	return nil
}

// Installation implement urlenc.URLValuesUnmarshaler.
func (m *Installation) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "installer":
				m.Installer = vals[0]
			case "spec":
				m.Spec = vals[0]
			}
		}
	}
	return nil
}

// Form implement urlenc.URLValuesUnmarshaler.
func (m *Form) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "title":
				m.Title = vals[0]
			case "headers":
				m.Headers = vals
			}
		}
	}
	return nil
}

// Line implement urlenc.URLValuesUnmarshaler.
func (m *Line) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "items":
				m.Items = vals
			}
		}
	}
	return nil
}

// ListOpusTypesRespData implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusTypesRespData) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "total":
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Total = uint32(val)
			}
		}
	}
	return nil
}

// CatalogInfo implement urlenc.URLValuesUnmarshaler.
func (m *CatalogInfo) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// OpusType implement urlenc.URLValuesUnmarshaler.
func (m *OpusType) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "type":
				m.Type = vals[0]
			case "name":
				m.Name = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			}
		}
	}
	return nil
}

// PutOnArtifactsReq implement urlenc.URLValuesUnmarshaler.
func (m *PutOnArtifactsReq) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgID":
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.OrgID = uint32(val)
			case "userID":
				m.UserID = vals[0]
			case "name":
				m.Name = vals[0]
			case "version":
				m.Version = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			case "summary":
				m.Summary = vals[0]
			case "labels":
				m.Labels = vals
			case "catalog":
				m.Catalog = vals[0]
			case "logoURL":
				m.LogoURL = vals[0]
			case "desc":
				m.Desc = vals[0]
			case "contactName":
				m.ContactName = vals[0]
			case "contactURL":
				m.ContactURL = vals[0]
			case "contactEmail":
				m.ContactEmail = vals[0]
			case "isOpenSourced":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsOpenSourced = val
			case "opensourceURL":
				m.OpensourceURL = vals[0]
			case "licenseName":
				m.LicenseName = vals[0]
			case "licenseURL":
				m.LicenseURL = vals[0]
			case "homepageName":
				m.HomepageName = vals[0]
			case "homepageURL":
				m.HomepageURL = vals[0]
			case "homepageLogoURL":
				m.HomepageLogoURL = vals[0]
			case "isDownloadable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDownloadable = val
			case "downloadURL":
				m.DownloadURL = vals[0]
			case "installation":
				if m.Installation == nil {
					m.Installation = &ArtifactsInstallation{}
				}
			case "installation.releaseID":
				if m.Installation == nil {
					m.Installation = &ArtifactsInstallation{}
				}
				m.Installation.ReleaseID = vals[0]
			}
		}
	}
	return nil
}

// ArtifactsInstallation implement urlenc.URLValuesUnmarshaler.
func (m *ArtifactsInstallation) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "releaseID":
				m.ReleaseID = vals[0]
			}
		}
	}
	return nil
}

// PutOffArtifactsReq implement urlenc.URLValuesUnmarshaler.
func (m *PutOffArtifactsReq) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgID":
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.OrgID = uint32(val)
			case "userID":
				m.UserID = vals[0]
			case "opusID":
				m.OpusID = vals[0]
			case "versionID":
				m.VersionID = vals[0]
			}
		}
	}
	return nil
}

// PutOnExtensionsReq implement urlenc.URLValuesUnmarshaler.
func (m *PutOnExtensionsReq) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgID":
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.OrgID = uint32(val)
			case "userID":
				m.UserID = vals[0]
			case "type":
				m.Type = vals[0]
			case "name":
				m.Name = vals[0]
			case "version":
				m.Version = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			case "displayNameI18n":
				m.DisplayNameI18N = vals[0]
			case "summary":
				m.Summary = vals[0]
			case "summaryI18n":
				m.SummaryI18N = vals[0]
			case "labels":
				m.Labels = vals
			case "catalog":
				m.Catalog = vals[0]
			case "logoURL":
				m.LogoURL = vals[0]
			case "level":
				m.Level = vals[0]
			case "mode":
				m.Mode = vals[0]
			case "desc":
				m.Desc = vals[0]
			case "descI18n":
				m.DescI18N = vals[0]
			case "contactName":
				m.ContactName = vals[0]
			case "contactURL":
				m.ContactURL = vals[0]
			case "contactEmail":
				m.ContactEmail = vals[0]
			case "isOpenSourced":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsOpenSourced = val
			case "opensourceURL":
				m.OpensourceURL = vals[0]
			case "licenseName":
				m.LicenseName = vals[0]
			case "licenseURL":
				m.LicenseURL = vals[0]
			case "homepageName":
				m.HomepageName = vals[0]
			case "homepageURL":
				m.HomepageURL = vals[0]
			case "homepageLogoURL":
				m.HomepageLogoURL = vals[0]
			case "isDownloadable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDownloadable = val
			case "downloadURL":
				m.DownloadURL = vals[0]
			case "i18n":
				m.I18N = vals[0]
			case "isDefault":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDefault = val
			}
		}
	}
	return nil
}

// ListOpusVersionsReq implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusVersionsReq) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "opusID":
				m.OpusID = vals[0]
			}
		}
	}
	return nil
}

// ListOpusVersionsResp implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusVersionsResp) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.Id = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Seconds = val
			case "data.createdAt.nanos":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Nanos = int32(val)
			case "data.updatedAt":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "data.updatedAt.seconds":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Seconds = val
			case "data.updatedAt.nanos":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Nanos = int32(val)
			case "data.orgID":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.OrgID = uint32(val)
			case "data.orgName":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.OrgName = vals[0]
			case "data.creatorID":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.CreatorID = vals[0]
			case "data.updaterID":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.UpdaterID = vals[0]
			case "data.level":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.Level = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.Type = vals[0]
			case "data.typeName":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.TypeName = vals[0]
			case "data.name":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.Name = vals[0]
			case "data.displayName":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.DisplayName = vals[0]
			case "data.catalog":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.Catalog = vals[0]
			case "data.defaultVersionID":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.DefaultVersionID = vals[0]
			case "data.latestVersionID":
				if m.Data == nil {
					m.Data = &ListOpusVersionsRespData{}
				}
				m.Data.LatestVersionID = vals[0]
			case "userIDs":
				m.UserIDs = vals
			}
		}
	}
	return nil
}

// ListOpusVersionsRespData implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusVersionsRespData) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "createdAt":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
			case "createdAt.seconds":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreatedAt.Seconds = val
			case "createdAt.nanos":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.CreatedAt.Nanos = int32(val)
			case "updatedAt":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "updatedAt.seconds":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdatedAt.Seconds = val
			case "updatedAt.nanos":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.UpdatedAt.Nanos = int32(val)
			case "orgID":
				val, err := strconv.ParseUint(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.OrgID = uint32(val)
			case "orgName":
				m.OrgName = vals[0]
			case "creatorID":
				m.CreatorID = vals[0]
			case "updaterID":
				m.UpdaterID = vals[0]
			case "level":
				m.Level = vals[0]
			case "type":
				m.Type = vals[0]
			case "typeName":
				m.TypeName = vals[0]
			case "name":
				m.Name = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			case "catalog":
				m.Catalog = vals[0]
			case "defaultVersionID":
				m.DefaultVersionID = vals[0]
			case "latestVersionID":
				m.LatestVersionID = vals[0]
			}
		}
	}
	return nil
}

// ListOpusVersionRespDataVersion implement urlenc.URLValuesUnmarshaler.
func (m *ListOpusVersionRespDataVersion) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "createdAt":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
			case "createdAt.seconds":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreatedAt.Seconds = val
			case "createdAt.nanos":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.CreatedAt.Nanos = int32(val)
			case "updatedAt":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "updatedAt.seconds":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdatedAt.Seconds = val
			case "updatedAt.nanos":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.UpdatedAt.Nanos = int32(val)
			case "creatorID":
				m.CreatorID = vals[0]
			case "updaterID":
				m.UpdaterID = vals[0]
			case "version":
				m.Version = vals[0]
			case "summary":
				m.Summary = vals[0]
			case "labels":
				m.Labels = vals
			case "logoURL":
				m.LogoURL = vals[0]
			case "isValid":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsValid = val
			case "ref":
				m.Ref = vals[0]
			case "desc":
				m.Desc = vals[0]
			case "contactName":
				m.ContactName = vals[0]
			case "contactURL":
				m.ContactURL = vals[0]
			case "contactEmail":
				m.ContactEmail = vals[0]
			case "isOpenSourced":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsOpenSourced = val
			case "opensourceURL":
				m.OpensourceURL = vals[0]
			case "licenceName":
				m.LicenceName = vals[0]
			case "licenceURL":
				m.LicenceURL = vals[0]
			case "homepageName":
				m.HomepageName = vals[0]
			case "homepageURL":
				m.HomepageURL = vals[0]
			case "homepageLogoURL":
				m.HomepageLogoURL = vals[0]
			case "isDownloadable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDownloadable = val
			case "downloadURL":
				m.DownloadURL = vals[0]
			case "readmeLang":
				m.ReadmeLang = vals[0]
			case "readmeLangName":
				m.ReadmeLangName = vals[0]
			case "readme":
				m.Readme = vals[0]
			}
		}
	}
	return nil
}

// PutOnOpusResp implement urlenc.URLValuesUnmarshaler.
func (m *PutOnOpusResp) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "opusID":
				m.OpusID = vals[0]
			case "versionID":
				m.VersionID = vals[0]
			}
		}
	}
	return nil
}

// I18N implement urlenc.URLValuesUnmarshaler.
func (m *I18N) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
