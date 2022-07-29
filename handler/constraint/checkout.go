package constraint

import (
	"context"
	"eureka/database"
	"fmt"
	"os"
)

func GetConstraintsAsGeojson(tableName string) (*map[string]interface{}, error) {
	dbstring := fmt.Sprintf(`
	SELECT jsonb_build_object(
        'type',
        'FeatureCollection',
        'features',
        jsonb_agg(features.feature)
    )
	FROM (
        SELECT jsonb_build_object(
                'type',
                'Feature',
                'geometry',
                ST_AsGeoJSON(geom)::jsonb,
                'properties',
                to_jsonb(inputs) - 'gid' - 'geom'
            ) AS feature
        FROM (
                SELECT *
                FROM public."%s"
                limit 10
            ) inputs
    ) features;
	`, tableName)
    
	// fmt.Println(dbstring)

	var geojson map[string]interface{}
	err := database.DB.QueryRow(context.Background(), dbstring).Scan(&geojson)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &geojson, nil
}
