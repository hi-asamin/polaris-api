package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/location"

	"polaris-api/domain"
	"polaris-api/interface/types"
)

type LocationServiceRepository struct{}

func (r *LocationServiceRepository) GeocodeAddress(address string) (*types.Geometry, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, domain.Wrap(err, 500, "unable to load AWS configuration")
	}

	// Create a new Amazon Location Service client
	locationSvc := location.NewFromConfig(cfg)

	// Call the Amazon Location Service to geocode the address
	input := &location.SearchPlaceIndexForTextInput{
		IndexName: aws.String("polaris-geocoder-index"),
		Text:      aws.String(address),
	}

	result, err := locationSvc.SearchPlaceIndexForText(context.TODO(), input)
	if err != nil {
		return nil, domain.Wrap(err, 500, "failed to search place index")
	}

	// Extract the latitude and longitude from the response
	if len(result.Results) == 0 {
		return nil, domain.New(404, "no results found for the address")
	}

	point := result.Results[0].Place.Geometry.Point
	return &types.Geometry{
		Latitude:  point[1],
		Longitude: point[0],
	}, nil
}
