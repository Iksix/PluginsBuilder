package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"plugins-builder/internal/config"
	"time"
)

func main() {
	config.Read()
	cfg := &config.Instance
	fmt.Println("Settings readed ✔")
	fmt.Println("Starting build...")
	cmd := exec.Command("dotnet", "build")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Plugin builded ✔")
	exist, err := exists("./build")
	if err != nil {
		fmt.Println(err)
	}
	if !exist {
		os.Mkdir("./build", os.ModeDir)
	}
	exist, err = exists(fmt.Sprintf("./build/%s", cfg.PluginName))
	if err != nil {
		fmt.Println(err)
	}
	if !exist {
		os.Mkdir(fmt.Sprintf("./build/%s", cfg.PluginName), os.ModeDir)
		os.MkdirAll(fmt.Sprintf("./build/%s/plugins/%s", cfg.PluginName, cfg.PluginName), os.ModeDir)
		if cfg.HasApi {
			os.MkdirAll(fmt.Sprintf("./build/%s/shared/%s", cfg.PluginName, cfg.ApiName), os.ModeDir)
		}
	}
	fmt.Println("Copy files to build dir...")
	copy(fmt.Sprintf("./bin/Debug/net8.0/%s.dll", cfg.PluginName), fmt.Sprintf("./build/%s/plugins/%s/%s.dll", cfg.PluginName, cfg.PluginName, cfg.PluginName))
	copy(fmt.Sprintf("./bin/Debug/net8.0/%s.pdb", cfg.PluginName), fmt.Sprintf("./build/%s/plugins/%s/%s.pdb", cfg.PluginName, cfg.PluginName, cfg.PluginName))
	if cfg.HasApi {
		copy(fmt.Sprintf("./bin/Debug/net8.0/%s.dll", cfg.ApiName), fmt.Sprintf("./build/%s/shared/%s/%s.dll", cfg.PluginName, cfg.ApiName, cfg.ApiName))
		copy(fmt.Sprintf("./bin/Debug/net8.0/%s.pdb", cfg.ApiName), fmt.Sprintf("./build/%s/shared/%s/%s.pdb", cfg.PluginName, cfg.ApiName, cfg.ApiName))
	}
	fmt.Println("NICE ✔")
	time.Sleep(time.Second * 2)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func copy(src string, dst string) {
	// Read all content of src to data, may cause OOM for a large file.
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	checkErr(err)
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
