package sql

func FindNearBySpots() string {
	return `
  SELECT 
    "Place".id AS pid,
    "Place".name,
    "Media".id AS mid, 
    "Media".media_url AS src,
    "Media".media_type AS type,
    "Media".alt_text AS alt,
    ST_Distance(
      "Place".geometry,
      ST_SetSRID(ST_MakePoint($1, $2), 4326)
    ) AS distance
  FROM 
    "Place"
  INNER JOIN
    "Media" 
    ON
      "Place".id = "Media".place_id
  WHERE
    "Place".geometry IS NOT NULL AND
    -- 1km以内にある場所
    ST_DWithin(
      "Place".geometry,
      ST_SetSRID(ST_MakePoint($1, $2), 4326),
      $3
    )
    AND "Place".id != $4 -- 指定された場所IDを除外
  ORDER BY
    distance ASC
  LIMIT $5;
  `
}
