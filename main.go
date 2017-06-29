package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

var confGroupSuffix = "_CONF_GROUP"

func main() {
	var envPrefix, targetFile string
	flag.StringVar(&envPrefix, "prefix", "", "variable prefix will be processed")
	flag.StringVar(&targetFile, "target", "", "variable prefix will be processed")

	flag.Parse()

	if envPrefix == "" || targetFile == "" {
		fmt.Println("prefix & target argument is required")
		os.Exit(1)
	}

	var trimPrefix string
	if strings.HasSuffix(envPrefix, "_") {
		trimPrefix = envPrefix
	} else {
		trimPrefix = fmt.Sprintf("%s_", envPrefix)
	}
	_env := os.Environ()
	content := ""
	filteredEnv := Filter(_env, func(s string) bool { return strings.HasPrefix(strings.Split(s, "=")[0], envPrefix) })
	trimmedEnv := Map(filteredEnv, func(s string) string { return strings.TrimPrefix(s, trimPrefix) })
	group := Filter(trimmedEnv, func(s string) bool { return strings.HasSuffix(strings.Split(s, "=")[0], confGroupSuffix) })
	attr := Filter(trimmedEnv, func(s string) bool { return !strings.HasSuffix(strings.Split(s, "=")[0], confGroupSuffix) })

	for _, g := range group {
		gsplit := strings.Split(g, "=")
		gkey := strings.TrimSuffix(gsplit[0], confGroupSuffix)
		gval := gsplit[1]
		content += fmt.Sprintf("\n[%s]\n", gval)
		gprefix := fmt.Sprintf("%s_", gkey)
		gattr := Filter(attr, func(s string) bool { return strings.HasPrefix(s, gprefix) })
		for _, a := range gattr {
			asplit := strings.Split(a, "=")
			akey := strings.TrimPrefix(asplit[0], gprefix)
			aval := asplit[1]
			content += fmt.Sprintf("%s = %s\n", strings.ToLower(akey), aval)
		}
	}

	ioutil.WriteFile(targetFile, []byte(content), 0644)
}
