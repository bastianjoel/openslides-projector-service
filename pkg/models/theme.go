package models

type Theme struct {
	Name                   string
	Warn200                *string
	Warn500                string
	Accent400              *string
	Accent50               *string
	Accent500              string
	Accent600              *string
	AccentA200             *string
	ID                     int
	Primary200             *string
	Primary300             *string
	Primary700             *string
	Yes                    *string
	WarnA700               *string
	Abstain                *string
	Accent800              *string
	Primary600             *string
	PrimaryA100            *string
	PrimaryA200            *string
	WarnA200               *string
	Primary400             *string
	ThemeForOrganizationID *int
	Warn50                 *string
	Warn700                *string
	WarnA100               *string
	PrimaryA700            *string
	Accent100              *string
	Accent900              *string
	AccentA700             *string
	Headbar                *string
	OrganizationID         int
	Warn300                *string
	Warn400                *string
	Warn800                *string
	Accent200              *string
	Accent300              *string
	Accent700              *string
	Primary100             *string
	Warn100                *string
	Warn900                *string
	Warn600                *string
	WarnA400               *string
	AccentA100             *string
	No                     *string
	Primary50              *string
	Primary500             string
	PrimaryA400            *string
	AccentA400             *string
	Primary800             *string
	Primary900             *string
}

func (m Theme) CollectionName() string {
	return "theme"
}
