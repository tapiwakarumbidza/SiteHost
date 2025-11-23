package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.recruitment.shq.nz"
	apiKey  = "h523hDtETbkJ3nSJL323hjYLXbCyDaRZ"
	clientID = "100"
)

type Domain struct {
	DomainID string `json:"domain_id"`
	Name     string `json:"name"`
	Zones    []Zone `json:"zones"`
}

type Zone struct {
	ZoneID string `json:"zone_id"`
	Name   string `json:"name"`
}

type DomainsResponse struct {
	Domains []Domain `json:"domains"`
}

type DNSRecord struct {
	RecordID string `json:"record_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
}

type ZoneRecordsResponse struct {
	Records []DNSRecord `json:"records"`
}

func fetchDomains() ([]Domain, error) {
	url := fmt.Sprintf("%s/domains/%s?api_key=%s", baseURL, clientID, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch domains: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var response DomainsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return response.Domains, nil
}

func fetchZoneRecords(zoneID string) ([]DNSRecord, error) {
	url := fmt.Sprintf("%s/zones/%s?api_key=%s", baseURL, zoneID, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch zone records: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var response ZoneRecordsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return response.Records, nil
}

func main() {
	fmt.Println("SiteHost Recruitment Challenge - Part One")
	fmt.Println("==========================================\n")

	// Fetch domains for Business Systems International (client_id=100)
	domains, err := fetchDomains()
	if err != nil {
		fmt.Printf("Error fetching domains: %v\n", err)
		return
	}

	fmt.Printf("Found %d domain(s) for Business Systems International (client_id=%s)\n\n", len(domains), clientID)

	for _, domain := range domains {
		fmt.Printf("Domain: %s (ID: %s)\n", domain.Name, domain.DomainID)

		if len(domain.Zones) == 0 {
			fmt.Println("  No DNS zones found for this domain")
		} else {
			fmt.Printf("  DNS Zones (%d):\n", len(domain.Zones))
			for _, zone := range domain.Zones {
				fmt.Printf("    - %s (ID: %s)\n", zone.Name, zone.ZoneID)

				// Fetch DNS records for this zone
				records, err := fetchZoneRecords(zone.ZoneID)
				if err != nil {
					fmt.Printf("      Error fetching records: %v\n", err)
					continue
				}

				if len(records) == 0 {
					fmt.Println("      No DNS records found")
				} else {
					fmt.Printf("      DNS Records (%d):\n", len(records))
					for _, record := range records {
						fmt.Printf("        [%s] %s -> %s (TTL: %d)\n", record.Type, record.Name, record.Content, record.TTL)
					}
				}
			}
		}
		fmt.Println()
	}
}
