package query

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"golang.org/x/exp/slices"
)

func getFeatureNames(data []map[string]string) []string {
	keys := make([]string, 0, len(data[0]))
	for k := range data[0] {
		keys = append(keys, k)
	}

	return keys
}

func ValidateParams(data []map[string]string, query url.Values) error {
	featureNames := getFeatureNames(data)

	for k := range query {
		if !slices.Contains(featureNames, k) && k != "limit" {
			return errors.New(fmt.Sprintf("Query parameter %s is not valid", k))
		}
	}

	_, err := getLimit(query)
	if err != nil {
		fmt.Println("fuckkk")
		return err
	}

	return nil
}

func getLimit(query url.Values) (int, error) {
	limitStr := query.Get("limit")
	if limitStr == "" {
		return 0, nil
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0, err
	}

	return limit, nil
}

func QueryData(query url.Values, data []map[string]string) ([]map[string]string, error) {
	limit, err := getLimit(query)
	if err != nil {
		return nil, err
	}

	results := make([]map[string]string, 0, len(data))
	for _, row := range data {
		satisfiesQuery := true
		for k := range query {
			if k != "limit" && row[k] != query.Get(k) {
				satisfiesQuery = false
			}
		}

		if satisfiesQuery && (len(results) < limit || limit == 0)  {
			results = append(results, row)
		}
	}
	return results, nil
}
