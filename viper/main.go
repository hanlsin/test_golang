package main

import (
	"flag"
	"fmt"

	"os"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("viper test")

	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	confYPath := flags.String("y", "test.yaml", "Set yaml config file")
	confJPath := flags.String("j", "test.json", "Set json config file")
	flags.Parse(os.Args[1:])

	fmt.Println("Yaml config file:", *confYPath)
	viper.SetConfigFile(*confYPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", viper.GetString("name"))

	fmt.Println("---------------------------------------------")
	viper.Reset()

	fmt.Println("Json config file:", *confJPath)
	viper.SetConfigFile(*confJPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Host:", viper.GetString("host"))
	fmt.Println("Host address:", viper.GetString("host.address"))
	fmt.Println("Host port:", viper.GetInt("host.port"))

	fmt.Println("---------------------------------------------")
	viper.Reset()

	viper.SetConfigName("test")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", viper.GetString("name"))
	fmt.Println("Host address:", viper.GetString("host.address"))
	fmt.Println("Host port:", viper.GetInt("host.port"))
}
