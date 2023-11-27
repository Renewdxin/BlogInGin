package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// The NewSetting function initializes the basic properties of the project
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
