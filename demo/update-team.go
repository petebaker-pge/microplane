package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var repoToOwner = map[string]string{
	"activate-students":                "eng-instant-login",
	"badge-generator-service":          "eng-instant-login",
	"badge-generator-worker":           "eng-instant-login",
	"badge-invalidation-worker":        "eng-instant-login",
	"broward-apps-csv":                 "eng-instant-login",
	"cil-event-logs-service":           "eng-instant-login",
	"cil-prov":                         "eng-instant-login",
	"cil-reliability-dashboard":        "eng-instant-login",
	"cil-vision":                       "eng-instant-login",
	"clever-com-router":                "eng-instant-login",
	"distribute-login-info":            "eng-instant-login",
	"district-authorizations-launcher": "eng-instant-login",
	"district-illerts-nightly":         "eng-instant-login",
	"district-user-service":            "eng-instant-login",
	"il-config-service":                "eng-instant-login",
	"il-loadtest-worker":               "eng-instant-login",
	"il-test-app":                      "eng-instant-login",
	"il-token-service":                 "eng-instant-login",
	"instant-login-api":                "eng-instant-login",
	"integration-tester":               "eng-instant-login",
	"ios-sdk":                          "eng-instant-login",
	"launchpad":                        "eng-instant-login",
	"lockbox-connectors":               "eng-instant-login",
	"lockbox-sftp":                     "eng-instant-login",
	"lockbox-sftp-watcher":             "eng-instant-login",
	"oauth":                            "eng-instant-login",
	"pdf-generator":                    "eng-instant-login",
	"resolve-ip":                       "eng-instant-login",
	"aeries":                           "eng-secure-sync",
	"api-router":                       "eng-secure-sync",
	"app-mailer":                       "eng-secure-sync",
	"app-service":                      "eng-secure-sync",
	"app-view-service":                 "eng-secure-sync",
	"apps-dashboard":                   "eng-secure-sync",
	"clever-cron":                      "eng-secure-sync",
	"clever-csharp":                    "eng-secure-sync",
	"clever-ga":                        "eng-secure-sync",
	"clever-go":                        "eng-secure-sync",
	"clever-java":                      "eng-secure-sync",
	"clever-js":                        "eng-secure-sync",
	"clever-php":                       "eng-secure-sync",
	"clever-python":                    "eng-secure-sync",
	"clever-ruby":                      "eng-secure-sync",
	"clever-sftp":                      "eng-secure-sync",
	"clever-sftp-watcher":              "eng-secure-sync",
	"cron-balancer":                    "eng-secure-sync",
	"csv-marshal-go":                   "eng-secure-sync",
	"csv-processor-2":                  "eng-secure-sync",
	"csv-upload":                       "eng-secure-sync",
	"custom-data-service":              "eng-secure-sync",
	"custom-sections-producer":         "eng-secure-sync",
	"dashboard-picker":                 "eng-secure-sync",
	"data-cleanse":                     "eng-secure-sync",
	"debouncer-service":                "eng-secure-sync",
	"debouncer-watcher":                "eng-secure-sync",
	"delete-district":                  "eng-secure-sync",
	"district-authorizations":          "eng-secure-sync",
	"district-data-validator":          "eng-secure-sync",
	"district-search-api":              "eng-secure-sync",
	"district-secure-share":            "eng-secure-sync",
	"district-view-service":            "eng-secure-sync",
	"docsmerger":                       "eng-secure-sync",
	"elasticsearch":                    "eng-secure-sync",
	"events-api":                       "eng-secure-sync",
	"events-ttl":                       "eng-secure-sync",
	"gaprov":                           "eng-secure-sync",
	"gaprov-scheduler":                 "eng-secure-sync",
	"gaprov-to-clever":                 "eng-secure-sync",
	"gaprov-worker":                    "eng-secure-sync",
	"gearman-admin":                    "eng-secure-sync",
	"gearman-admin-load-tester":        "eng-secure-sync",
	"gearman-load-logger":              "eng-secure-sync",
	"ggprov-scheduler":                 "eng-secure-sync",
	"ggprov-worker":                    "eng-secure-sync",
	"google-groups-prov":               "eng-secure-sync",
	"http-science":                     "eng-secure-sync",
	"hubble":                           "eng-secure-sync",
	"ic-normalizer":                    "eng-secure-sync",
	"ic-oneroster-api":                 "eng-secure-sync",
	"illuminate":                       "eng-secure-sync",
	"infinitecampus":                   "eng-secure-sync",
	"inow":                             "eng-secure-sync",
	"jijiprov":                         "eng-secure-sync",
	"json-processor":                   "eng-secure-sync",
	"lausd-csv-normalizer":             "eng-secure-sync",
	"leakybucket":                      "eng-secure-sync",
	"legacy-normalizer":                "eng-secure-sync",
	"legacy-schools-dashboard":         "eng-secure-sync",
	"matchmaker-service":               "eng-secure-sync",
	"matchmaker-spawner":               "eng-secure-sync",
	"matchmaker-worker":                "eng-secure-sync",
	"miami":                            "eng-secure-sync",
	"mongo-op-throttler":               "eng-secure-sync",
	"mongo-system-copier":              "eng-secure-sync",
	"mp-finalizer":                     "eng-secure-sync",
	"mp-locker":                        "eng-secure-sync",
	"mp-trigger":                       "eng-secure-sync",
	"mp-workflow-service":              "eng-secure-sync",
	"mpl-elasticsearch-sis":            "eng-secure-sync",
	"mpl-events":                       "eng-secure-sync",
	"mpl-sis":                          "eng-secure-sync",
	"mpl-summaries":                    "eng-secure-sync",
	"mpt-app-sharing":                  "eng-secure-sync",
	"mpt-data-gator":                   "eng-secure-sync",
	"mpt-dead-links":                   "eng-secure-sync",
	"mpt-district-sharing":             "eng-secure-sync",
	"mpt-metadata":                     "eng-secure-sync",
	"mpt-scopes":                       "eng-secure-sync",
	"mpt-view-differ":                  "eng-secure-sync",
	"nagbot":                           "eng-secure-sync",
	"oneroster-csv-upload":             "eng-secure-sync",
	"oneroster-normalizer":             "eng-secure-sync",
	"oneroster11-normalizer":           "eng-secure-sync",
	"oplog-dump":                       "eng-secure-sync",
	"permissions-service":              "eng-secure-sync",
	"pipeline-combinator":              "eng-secure-sync",
	"pipeline-combinator-trigger":      "eng-secure-sync",
	"pipeline-data-quality":            "eng-secure-sync",
	"pipeline-finalizer":               "eng-secure-sync",
	"pipeline-normalizer":              "eng-secure-sync",
	"pipeline-rollup":                  "eng-secure-sync",
	"powerschool":                      "eng-secure-sync",
	"powerschool2":                     "eng-secure-sync",
	"ps-emails":                        "eng-secure-sync",
	"qa":                               "eng-secure-sync",
	"quest":                            "eng-secure-sync",
	"redirector-service":               "eng-secure-sync",
	"reverse-sftp-mapping":             "eng-secure-sync",
	"sandbox-events":                   "eng-secure-sync",
	"school-admin-upload":              "eng-secure-sync",
	"school-mdr-worker":                "eng-secure-sync",
	"schoolinsight":                    "eng-secure-sync",
	"scope-service":                    "eng-secure-sync",
	"sd2":                              "eng-secure-sync",
	"service-performance-metrics": "eng-secure-sync",
	"sftp-sync-router":            "eng-secure-sync",
	"sharing-mailer":              "eng-secure-sync",
	"signal-me-maybe":             "eng-secure-sync",
	"sis-api":                     "eng-secure-sync",
	"skyward":                     "eng-secure-sync",
	"skyward-api":                 "eng-secure-sync",
	"skyward-api-normalizer":      "eng-secure-sync",
	"skyward-normalizer":          "eng-secure-sync",
	"space":                       "eng-secure-sync",
	"spec-service":                "eng-secure-sync",
	"sphinx":                      "eng-secure-sync",
	"sphinx-api":                  "eng-secure-sync",
	"system-converter":            "eng-secure-sync",
	"systemic":                    "eng-secure-sync",
	"tag-middleware":              "eng-secure-sync",
	"titanium-transactions":       "eng-secure-sync",
	"user-search-service":         "eng-secure-sync",
	"user-search-ui":              "eng-secure-sync",
	"user-search-updater":         "eng-secure-sync",
}

func main() {
	output, err := exec.Command("git", "remote", "get-url", "--push", "origin").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	remoteURL := strings.TrimSpace(string(output))
	remoteURLSplit := strings.Split(string(remoteURL), "/")
	repoName := strings.Replace(remoteURLSplit[len(remoteURLSplit)-1], ".git", "", 1)
	newOwner, ok := repoToOwner[repoName]
	if !ok {
		os.Exit(0)
	}

	// modify all files launch/
	files, err := ioutil.ReadDir("./launch")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if output, err := exec.Command("sed", "-ibak", "-e",
			fmt.Sprintf("s/team:.*/team: %s/", newOwner),
			fmt.Sprintf("launch/%s", file.Name()),
		).CombinedOutput(); err != nil {
			log.Fatal(string(output))
		}
		if output, err := exec.Command("rm", fmt.Sprintf("launch/%sbak", file.Name())).CombinedOutput(); err != nil {
			log.Fatal(string(output))
		}
	}
}
