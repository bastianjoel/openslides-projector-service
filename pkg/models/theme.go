package models

type Theme struct {
	Abstain                *string
	Accent100              *string
	Accent200              *string
	Accent300              *string
	Accent400              *string
	Accent50               *string
	Accent500              string
	Accent600              *string
	Accent700              *string
	Accent800              *string
	Accent900              *string
	AccentA100             *string
	AccentA200             *string
	AccentA400             *string
	AccentA700             *string
	Headbar                *string
	ID                     int
	Name                   string
	No                     *string
	OrganizationID         int
	Primary100             *string
	Primary200             *string
	Primary300             *string
	Primary400             *string
	Primary50              *string
	Primary500             string
	Primary600             *string
	Primary700             *string
	Primary800             *string
	Primary900             *string
	PrimaryA100            *string
	PrimaryA200            *string
	PrimaryA400            *string
	PrimaryA700            *string
	ThemeForOrganizationID *int
	Warn100                *string
	Warn200                *string
	Warn300                *string
	Warn400                *string
	Warn50                 *string
	Warn500                string
	Warn600                *string
	Warn700                *string
	Warn800                *string
	Warn900                *string
	WarnA100               *string
	WarnA200               *string
	WarnA400               *string
	WarnA700               *string
	Yes                    *string
}

func (m Theme) CollectionName() string {
	return "theme"
}
