package main

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"log"
	"os"
	"strconv"
)

func main() {
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	if len(os.Args) > 1 {
		zoneName := os.Args[1]
		zoneID, err := getZoneIDByName(api, zoneName)
		if err != nil {
			log.Fatal(err)
		}
		println("Zone name: "+zoneName, "\nZone ID: "+zoneID)
		records, err := getDNSRecords(ctx, api, zoneID)
		for _, r := range records {
			fmt.Printf("%s: %s\n", r.Name, r.Content)
		}

	} else {
		fmt.Println("Usage: go-cloudflare-info <zoneName>")
		fmt.Println("No domain name provided - listing all zones")
		zones, err := getZones(ctx, api)
		if err != nil {
			log.Fatal(err)
		}

		i := 1
		for _, z := range zones {
			fmt.Println(strconv.Itoa(i) + ". " + "ID: " + z.ID + " Domain: " + z.Name)
			i += 1
		}
	}
}

func getZones(ctx context.Context, api *cloudflare.API) ([]cloudflare.Zone, error) {
	zones, err := api.ListZones(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return zones, err
}

func getZoneIDByName(api *cloudflare.API, zoneName string) (string, error) {
	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
		log.Fatal(err)
	}

	return zoneID, err
}

func getDNSRecords(ctx context.Context, api *cloudflare.API, zoneID string) ([]cloudflare.DNSRecord, error) {
	recs, _, err := api.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		log.Fatal(err)
	}

	return recs, err
}
