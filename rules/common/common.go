package common

import "github.com/olebedev/when/rules"

var All = []rules.Rule{
	ISODate(rules.Override),
	SlashDMY(rules.Override),
}
