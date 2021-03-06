package scaleway

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	RdbWaitForTimeout = 10 * time.Minute
)

// rdbAPIWithRegion returns a new lb API and the region for a Create request
func rdbAPIWithRegion(d *schema.ResourceData, m interface{}) (*rdb.API, scw.Region, error) {
	meta := m.(*Meta)
	rdbApi := rdb.NewAPI(meta.scwClient)

	region, err := extractRegion(d, meta)
	return rdbApi, region, err
}

// rdbAPIWithRegionAndID returns an lb API with region and ID extracted from the state
func rdbAPIWithRegionAndID(m interface{}, id string) (*rdb.API, scw.Region, string, error) {
	meta := m.(*Meta)
	rdbApi := rdb.NewAPI(meta.scwClient)

	region, ID, err := parseRegionalID(id)
	return rdbApi, region, ID, err
}

func flattenRdbInstanceReadReplicas(readReplicas []*rdb.Endpoint) interface{} {
	replicasI := []map[string]interface{}(nil)
	for _, readReplica := range readReplicas {
		replicasI = append(replicasI, map[string]interface{}{
			"ip":   flattenIPPtr(readReplica.IP),
			"port": int(readReplica.Port),
			"name": flattenStringPtr(readReplica.Name),
		})
	}
	return replicasI
}
