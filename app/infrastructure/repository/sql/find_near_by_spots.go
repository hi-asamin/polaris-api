package sql

func FindNearBySpots() string {
	return `
  SELECT 
    "Place".id AS pid,
    "Place".name,
    "Place".state,
    "Place".city,
    "Media".id AS mid, 
    "Media".media_url AS src,
    "Media".media_type AS type,
    "Media".alt_text AS alt,
    CEIL(ST_Distance(
      ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
      "Place".geometry::geography
    ) / 1000) AS distance
  FROM 
    "Place"
  INNER JOIN
    "Media" 
    ON
      "Place".id = "Media".place_id
  WHERE
    "Place".geometry IS NOT NULL AND
    -- 指定範囲内にある場所
    ST_DWithin(
      ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
      "Place".geometry::geography,
      $3
    )
    AND "Place".id != $4 -- 指定された場所IDを除外
  ORDER BY
    distance ASC
  LIMIT $5;
  `
}
