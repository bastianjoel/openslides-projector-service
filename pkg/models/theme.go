package models

type Theme struct {
	Abstain                *string `json:"abstain"`
	Accent100              *string `json:"accent_100"`
	Accent200              *string `json:"accent_200"`
	Accent300              *string `json:"accent_300"`
	Accent400              *string `json:"accent_400"`
	Accent50               *string `json:"accent_50"`
	Accent500              string  `json:"accent_500"`
	Accent600              *string `json:"accent_600"`
	Accent700              *string `json:"accent_700"`
	Accent800              *string `json:"accent_800"`
	Accent900              *string `json:"accent_900"`
	AccentA100             *string `json:"accent_a100"`
	AccentA200             *string `json:"accent_a200"`
	AccentA400             *string `json:"accent_a400"`
	AccentA700             *string `json:"accent_a700"`
	Headbar                *string `json:"headbar"`
	ID                     int     `json:"id"`
	Name                   string  `json:"name"`
	No                     *string `json:"no"`
	OrganizationID         int     `json:"organization_id"`
	Primary100             *string `json:"primary_100"`
	Primary200             *string `json:"primary_200"`
	Primary300             *string `json:"primary_300"`
	Primary400             *string `json:"primary_400"`
	Primary50              *string `json:"primary_50"`
	Primary500             string  `json:"primary_500"`
	Primary600             *string `json:"primary_600"`
	Primary700             *string `json:"primary_700"`
	Primary800             *string `json:"primary_800"`
	Primary900             *string `json:"primary_900"`
	PrimaryA100            *string `json:"primary_a100"`
	PrimaryA200            *string `json:"primary_a200"`
	PrimaryA400            *string `json:"primary_a400"`
	PrimaryA700            *string `json:"primary_a700"`
	ThemeForOrganizationID *int    `json:"theme_for_organization_id"`
	Warn100                *string `json:"warn_100"`
	Warn200                *string `json:"warn_200"`
	Warn300                *string `json:"warn_300"`
	Warn400                *string `json:"warn_400"`
	Warn50                 *string `json:"warn_50"`
	Warn500                string  `json:"warn_500"`
	Warn600                *string `json:"warn_600"`
	Warn700                *string `json:"warn_700"`
	Warn800                *string `json:"warn_800"`
	Warn900                *string `json:"warn_900"`
	WarnA100               *string `json:"warn_a100"`
	WarnA200               *string `json:"warn_a200"`
	WarnA400               *string `json:"warn_a400"`
	WarnA700               *string `json:"warn_a700"`
	Yes                    *string `json:"yes"`
}

func (m Theme) CollectionName() string {
	return "theme"
}

func (m Theme) Get(field string) interface{} {
	switch field {
	case "abstain":
		return m.Abstain
	case "accent_100":
		return m.Accent100
	case "accent_200":
		return m.Accent200
	case "accent_300":
		return m.Accent300
	case "accent_400":
		return m.Accent400
	case "accent_50":
		return m.Accent50
	case "accent_500":
		return m.Accent500
	case "accent_600":
		return m.Accent600
	case "accent_700":
		return m.Accent700
	case "accent_800":
		return m.Accent800
	case "accent_900":
		return m.Accent900
	case "accent_a100":
		return m.AccentA100
	case "accent_a200":
		return m.AccentA200
	case "accent_a400":
		return m.AccentA400
	case "accent_a700":
		return m.AccentA700
	case "headbar":
		return m.Headbar
	case "id":
		return m.ID
	case "name":
		return m.Name
	case "no":
		return m.No
	case "organization_id":
		return m.OrganizationID
	case "primary_100":
		return m.Primary100
	case "primary_200":
		return m.Primary200
	case "primary_300":
		return m.Primary300
	case "primary_400":
		return m.Primary400
	case "primary_50":
		return m.Primary50
	case "primary_500":
		return m.Primary500
	case "primary_600":
		return m.Primary600
	case "primary_700":
		return m.Primary700
	case "primary_800":
		return m.Primary800
	case "primary_900":
		return m.Primary900
	case "primary_a100":
		return m.PrimaryA100
	case "primary_a200":
		return m.PrimaryA200
	case "primary_a400":
		return m.PrimaryA400
	case "primary_a700":
		return m.PrimaryA700
	case "theme_for_organization_id":
		return m.ThemeForOrganizationID
	case "warn_100":
		return m.Warn100
	case "warn_200":
		return m.Warn200
	case "warn_300":
		return m.Warn300
	case "warn_400":
		return m.Warn400
	case "warn_50":
		return m.Warn50
	case "warn_500":
		return m.Warn500
	case "warn_600":
		return m.Warn600
	case "warn_700":
		return m.Warn700
	case "warn_800":
		return m.Warn800
	case "warn_900":
		return m.Warn900
	case "warn_a100":
		return m.WarnA100
	case "warn_a200":
		return m.WarnA200
	case "warn_a400":
		return m.WarnA400
	case "warn_a700":
		return m.WarnA700
	case "yes":
		return m.Yes
	}

	return nil
}
