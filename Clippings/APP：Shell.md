---
title: "Shell"
source: "https://nilesoft.org/docs/configuration/settings"
author:
  - "[[Nilesoft]]"
published:
created: 2024-11-06
description: "Shell is a powerful context menu customizer with highly responsive for Windows File Explorer."
tags:
  - "clippings"
---
#### Settings

Settings are containers for storing default values.

```shell
settings
{
	// show menu delay value from 0 to 4000
	showdelay = 200

	// Prevent interaction with these windows or processes
	exclude
	{
		where = boolean value
		window = window name
		process = process name
	}

	tip = true
	// or
	tip
	{
		enabled = true

		// normal = [background, text]
		normal = [default, default]

		// normal = [background, text]
		primary = [#000, #fff]

		// info = [background, text]
		info = [#88f, #fff]

		// success = [background, text]
		success = [#8f8, #fff]

		// warning = [background, text]
		warning = [#ff8, #fff]

		// danger = [background, text]
		danger = [#f88, #fff]

		// max width value from 200 to 2000
		width = 400

		// opacity value from 0 to 100
		opacity = 100

		// radius size value from 0 to 3
		radius = 1

		time = 1.5

		padding = [8, 4]
	}

	// Disable/Enable modify items processing
	modify
	{
		enabled = boolean value
		image = [0 = disable, 1 = enable, 2 = auto reimage]

		// Allow/disallow modification of title
		title = boolean value

		// Allow/disallow modification of visibility
		visibility = boolean value

		// Allow/disallow modification of parent
		parent = boolean value

		// Allow/disallow modification of position
		position = boolean value

		// Allow/disallow to add separator
		separator = boolean value

		// auto set image and group
		auto = boolean value
	}

	// Disable/Enable new items processing
	new
	{
		enabled = boolean value
		// disable/enable image
		image = boolean value
	}
}
```

---