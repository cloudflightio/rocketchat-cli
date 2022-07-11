module github.com/cloudflightio/rocketchat-cli

go 1.18

require (
	github.com/RocketChat/Rocket.Chat.Go.SDK v0.0.0-20220708192748-b73dcb041214
	github.com/magiconair/properties v1.8.6
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/subosito/gotenv v1.4.0 // indirect
	golang.org/x/sys v0.0.0-20220708085239-5a0f0661e09d // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//hotfix till https://github.com/RocketChat/Rocket.Chat.Go.SDK/pull/52 is merged
replace github.com/RocketChat/Rocket.Chat.Go.SDK => github.com/evanofslack/Rocket.Chat.Go.SDK v0.0.0-20220615231022-f62f5cba832c
