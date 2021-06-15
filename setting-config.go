package main

import "github.com/spf13/viper"

func LoadSettingConfig() (config SettingConfig, err error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("setting")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type SettingConfig struct {
	Type                string `mapstructure:"type"`
	Mode                int    `mapstructure:"mode"`
	Number              int    `mapstructure:"number"`
	Difficulty          int    `mapstructure:"difficulty"`
	StaminaRecoveryTime int    `mapstructure:"staminaRecoveryTime"`
}

/*
{
    "type": "freeActivity Or freeBoss Or repalay Or Auto",  (free 打免費共鬥房間 如果是活動 就會參考number And difficuly  , auto 是有體力打指定 沒體力打共鬥)
    "mode": 1, (fre 模式才有效果, 決定要直接開始1 或者 開放等待2 或是 不開放等待3)
    "number": 1, (選關卡 由上而下)
    "difficulty":1,  (選難度 由上而下)
    "staminaRecoveryTime":7200 (auto 時才有效果 沒體力後過多久再次嘗試)
}
*/
