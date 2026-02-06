package meta

import (
	package meta

	import (
		"context"

		"github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs/pkg/ctxkeys"

		"github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs/pkg/config"
	)

	func SiteFromCtx(ctx context.Context) config.SiteConfig {
		if cfg, ok := ctx.Value(ctxkeys.SiteConfig).(config.SiteConfig); ok {
			return cfg
		}
		return config.SiteConfig{Name: "Ron's Designs and Signs"}
	}

	func SiteNameFromCtx(ctx context.Context) string {
		return SiteFromCtx(ctx).Name
	}

	func SiteURLFromCtx(ctx context.Context) string {
		return SiteFromCtx(ctx).URL
	}
