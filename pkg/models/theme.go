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
