package opoptions

import (
	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/util"
)

// WithFilterType allows for changing of the default filter type (subtree) for NETCONF operations.
func WithFilterType(s string) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*netconf.OperationOptions)

		if ok {
			c.FilterType = s

			return nil
		}

		return util.ErrIgnoredOption
	}
}

// WithDefaultType allows for changing of the default "default type" type for NETCONF operations.
func WithDefaultType(s string) util.Option {
	return func(o interface{}) error {
		c, ok := o.(*netconf.OperationOptions)

		if ok {
			c.DefaultType = s

			return nil
		}

		return util.ErrIgnoredOption
	}
}