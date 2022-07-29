package handler

import (
	"context"
	"eureka/database"
	"eureka/model"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAdvisory(c *fiber.Ctx) error {
	//Parse Body
	AdvisoryQuery_Input := new(model.AdvisoryQueryModel)
	if err := c.BodyParser(AdvisoryQuery_Input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body"})
	}

	//Validate Body
	err := model.ValidateAdvisoryQueryModel(AdvisoryQuery_Input)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error()})
	}

	//Run Layerset
	response := *RunAdvisoryLayerSet(AdvisoryQuery_Input)

	//Format response
	return c.JSON(response)
}

func RunAdvisoryLayerSet(input *model.AdvisoryQueryModel) *model.AdvisoryResponseModel {

	var output model.AdvisoryResponseModel

	RunLayerSets(input, &output)

	return &output
}

func RunLayerSets(input *model.AdvisoryQueryModel, output *model.AdvisoryResponseModel) {
	var waitGroup sync.WaitGroup
	//# of thread group
	waitGroup.Add(len(*&input.Layers))
	//Running all layers
	for _, cur := range input.Layers {
		curLayer := new(model.AdvisoryLayer)
		{
			curLayer.Name = cur
		}
		go RunLayer(&waitGroup, curLayer, input, output)
	}
	waitGroup.Wait()
}

func RunLayer(waitGroup *sync.WaitGroup, curLayer *model.AdvisoryLayer, input *model.AdvisoryQueryModel, output *model.AdvisoryResponseModel) {
	curTime := time.Now()
	switch curLayer.Name {
	default:
		curLayer.Details, _ = GetASGeojson(curLayer.Name, input.Geometry)
	}
	curLayer.QueryDuration = int(time.Since(curTime).Milliseconds())
	output.RuleSet = append(output.RuleSet, curLayer)
	//Lift Sync Lock
	defer waitGroup.Done()
}

func GetASGeojson(layer, location string) (*map[string]interface{}, error) {
	fetch_string := fmt.Sprintf(`
	SELECT *
	FROM public."%s"
	where 
	ST_Intersects(st_geomfromtext('%s'),
	geom)
	limit 10
	`, layer, location)

	var geojson map[string]interface{}
	err := database.DB.QueryRow(context.Background(), jsonb_build_object(fetch_string)).Scan(&geojson)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return &geojson, nil
}

func jsonb_build_object(original string) string {
	return fmt.Sprintf(`
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
                %s
        ) inputs
    ) features;
	`, original)
}
