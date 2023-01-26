package main

import (
	"fmt"
	"github.com/caseymrm/menuet"
	"github.com/maxrichie5/mac-clock-toolbar/internal/clocks"
	"os"
	"time"
)

func main() {
	//cmd := `"for profile in $(cat ~/.aws/credentials | grep '\[' | column -t -s '[]'); do for cluster in $(aws eks list-clusters --profile $profile --output text | awk '{ print $2 }'); do aws eks update-kubeconfig --name $cluster --profile $profile --alias $(echo $(echo $cluster | sed s/kount-//g | sed s/apps-cluster-//g | sed s/gitlab-cluster-//g)-$(column -t -s "-" <<<"$profile" | awk '{print $NF}')) ; done ; done"`
	//b, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(b))
	go doTheMenuThing()
	menuet.App().RunApplication()
	fmt.Println("Running")
}

func fatal(err error, msg string, args ...interface{}) {
	if err != nil {
		msg += ": %v"
		args = append(args, err)
	}
	fmt.Fprintf(os.Stderr, msg, args...)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		os.Exit(1)
	} else {
		// nil error means it was CLI usage issue
		fmt.Fprintf(os.Stderr, "Try '%s --help' for more details.\n", os.Args[0])
		os.Exit(2)
	}
}

//func buildMenuWithContexts(config *api.Config) []menuet.MenuItem {
//	contexts := getContextNames(config)
//	contexts = menu.FilterBadContexts(contexts)
//	contexts = menu.FilterOldContexts(contexts)
//	grouped := menu.GroupContextsByRole(contexts)
//	mu := menu.TurnContextsIntoMenu(config, grouped, menuThing)
//	return menu.AddSeparators(mu)
//}

func setMenuState() {
	menuet.App().Children = func() []menuet.MenuItem {
		return clocks.GetAllClocks()
	}
	menuet.App().SetMenuState(&menuet.MenuState{
		Title: clocks.GetActiveClocks(),
	})
	menuet.App().Label = "Mac Clock Toolbar"
}

func doTheMenuThing() {
	for {
		setMenuState()
		time.Sleep(time.Second)
	}
}

//func getContextNames(config *api.Config) []string {
//	keys := make([]string, len(config.Contexts))
//	i := 0
//	for k := range config.Contexts {
//		keys[i] = k
//		i++
//	}
//	return keys
//}
