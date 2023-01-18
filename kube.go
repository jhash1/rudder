package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/johnhashemiii/.kube/config", "location of your Kubeconfig file")
	flag.Parse()
	helmActionConfig := new(action.Configuration)

	err := helmActionConfig.Init(kube.GetConfig(*kubeconfig, "", ""), "", "", log.Printf)
	if err != nil {
		log.Fatal(err)
	}
	helmReleases, err := action.NewList(helmActionConfig).Run()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Showing Helm Releases:")
	for _, release := range helmReleases {
		fmt.Println("Chart Name:", release.Chart.Metadata.Name, "\n", "Release Name:", "\n", release.Name, "\n", "Installed: \n", release.Chart.AppVersion(), "\n",
			"Kubernetes Kubernetes Version Compatibility: \n", release.Chart.Metadata.KubeVersion, "\n",
			"Deprecated: \n", release.Chart.Metadata, release.Chart.Metadata.)
		for _, repoURL := range release.Chart.Metadata.Sources {
			fmt.Println(repoURL)
		}
	}

	type artifactHub struct {
		Packages []struct {
			PackageID                    string `json:"package_id"`
			Name                         string `json:"name"`
			NormalizedName               string `json:"normalized_name"`
			Stars                        int    `json:"stars"`
			Description                  string `json:"description"`
			Version                      string `json:"version"`
			AppVersion                   string `json:"app_version,omitempty"`
			Deprecated                   bool   `json:"deprecated"`
			Signed                       bool   `json:"signed"`
			ProductionOrganizationsCount int    `json:"production_organizations_count"`
			Ts                           int    `json:"ts"`
			Repository                   struct {
				URL                     string `json:"url"`
				Kind                    int    `json:"kind"`
				Name                    string `json:"name"`
				Official                bool   `json:"official"`
				DisplayName             string `json:"display_name"`
				RepositoryID            string `json:"repository_id"`
				ScannerDisabled         bool   `json:"scanner_disabled"`
				OrganizationName        string `json:"organization_name"`
				VerifiedPublisher       bool   `json:"verified_publisher"`
				OrganizationDisplayName string `json:"organization_display_name"`
			} `json:"repository"`
			SecurityReportSummary struct {
				Low      int `json:"low"`
				High     int `json:"high"`
				Medium   int `json:"medium"`
				Unknown  int `json:"unknown"`
				Critical int `json:"critical"`
			} `json:"security_report_summary,omitempty"`
			AllContainersImagesWhitelisted bool `json:"all_containers_images_whitelisted,omitempty"`
		} `json:"packages"`
	}
	// https://artifacthub.io/docs/api/#/Packages/getHelmPackageDetails
	// include "package_id": "8920517c-68de-4dbe-b29d-a2a778e872f9",
	//  "name": "nginx-ingress",
	// "security_report_summary": {
	// 	"low": 93,
	// 	"high": 15,
	// 	"medium": 32,
	// 	"unknown": 0,
	// 	"critical": 4
	//   },
	// "contains_security_updates": false, need to explore to justify upgrade
	//does include container images in a slice which can be used to check images seperately which
	//could contain a vuln which you need to contact maintainer of the chart.

	// create a method for printing out security findings for the artifact hub struct

	resp, err := http.Get("https://artifacthub.io/api/v1/packages/search?Helm?")
	if err != nil {
		log.Fatal(err)
	}

	artifactHub1 := artifactHub{}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}
	// fmt.Println(string(body))
	b := json.Unmarshal(body, &artifactHub1)
	if b != nil {
		fmt.Println("error:", err)
		fmt.Println(artifactHub1.Packages[0].Name)

		// func (ah *artifacthub) printSecuritySummary() {
		// 	for s,n := range *artifactHub{

	}
}


