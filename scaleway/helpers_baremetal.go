package scaleway

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	baremetal "github.com/scaleway/scaleway-sdk-go/api/baremetal/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	baremetalServerWaitForTimeout   = 60 * time.Minute
	baremetalServerRetryFuncTimeout = baremetalServerWaitForTimeout + time.Minute // some RetryFunc are calling a WaitFor
)

var baremetalServerResourceTimeout = baremetalServerRetryFuncTimeout + time.Minute

// instanceAPIWithZone returns a new baremetal API and the zone for a Create request
func baremetalAPIWithZone(d *schema.ResourceData, m interface{}) (*baremetal.API, scw.Zone, error) {
	meta := m.(*Meta)
	baremetalAPI := baremetal.NewAPI(meta.scwClient)

	zone, err := extractZone(d, meta)
	return baremetalAPI, zone, err
}

// instanceAPIWithZoneAndID returns an baremetal API with zone and ID extracted from the state
func baremetalAPIWithZoneAndID(m interface{}, id string) (*baremetal.API, scw.Zone, string, error) {
	meta := m.(*Meta)
	baremetalAPI := baremetal.NewAPI(meta.scwClient)

	zone, ID, err := parseZonedID(id)
	return baremetalAPI, zone, ID, err
}

// TODO: Remove it when SDK will handle it.
// baremetalOfferByName call baremetal API to get an offer by its exact name.
func baremetalOfferByName(baremetalAPI *baremetal.API, zone scw.Zone, offerName string) (*baremetal.Offer, error) {
	offerRes, err := baremetalAPI.ListOffers(&baremetal.ListOffersRequest{
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return nil, err
	}

	offerName = strings.ToUpper(offerName)
	for _, offer := range offerRes.Offers {
		if offer.Name == offerName {
			return offer, nil
		}
	}
	return nil, fmt.Errorf("cannot find the offer %s", offerName)
}

// TODO: Remove it when SDK will handle it.
// baremetalOfferByID call baremetal API to get an offer by its exact name.
func baremetalOfferByID(baremetalAPI *baremetal.API, zone scw.Zone, offerID string) (*baremetal.Offer, error) {
	offerRes, err := baremetalAPI.ListOffers(&baremetal.ListOffersRequest{
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		return nil, err
	}

	for _, offer := range offerRes.Offers {
		if offer.ID == offerID {
			return offer, nil
		}
	}
	return nil, fmt.Errorf("cannot find the offer %s", offerID)
}

func flattenBaremetalCPUs(cpus []*baremetal.CPU) interface{} {
	if cpus == nil {
		return nil
	}
	flattenedCPUs := []map[string]interface{}(nil)
	for _, cpu := range cpus {
		flattenedCPUs = append(flattenedCPUs, map[string]interface{}{
			"name":         cpu.Name,
			"core_count":   cpu.Cores,
			"frequency":    cpu.Frequency,
			"thread_count": cpu.Threads,
		})
	}
	return flattenedCPUs
}

func flattenBaremetalDisks(disks []*baremetal.Disk) interface{} {
	if disks == nil {
		return nil
	}
	flattenedDisks := []map[string]interface{}(nil)
	for _, disk := range disks {
		flattenedDisks = append(flattenedDisks, map[string]interface{}{
			"type":     disk.Type,
			"capacity": disk.Capacity,
		})
	}
	return flattenedDisks
}

func flattenBaremetalMemory(memories []*baremetal.Memory) interface{} {
	if memories == nil {
		return nil
	}
	flattenedMemories := []map[string]interface{}(nil)
	for _, memory := range memories {
		flattenedMemories = append(flattenedMemories, map[string]interface{}{
			"type":      memory.Type,
			"capacity":  memory.Capacity,
			"frequency": memory.Frequency,
			"ecc":       memory.Ecc,
		})
	}
	return flattenedMemories
}

func flattenBaremetalIPs(ips []*baremetal.IP) interface{} {
	if ips == nil {
		return nil
	}
	flattendIPs := []map[string]interface{}(nil)
	for _, ip := range ips {
		flattendIPs = append(flattendIPs, map[string]interface{}{
			"id":      ip.ID,
			"address": ip.Address,
			"reverse": ip.Reverse,
		})
	}
	return flattendIPs
}
