WITH places_with_distance AS (
  SELECT 
    "Place".id AS placeId,
    "Place".name,
    "Media".id AS mediaId, 
    "Media".media_url AS src,
    "Media".media_type AS type,
    "Media".alt_text AS alt,
    ST_Distance(
      "Place".geometry, 
      ST_SetSRID(ST_MakePoint($1, $2), 4326)
    ) AS distance
  FROM
    "Place"
  LEFT JOIN
    "Media" 
    ON
      "Place".id = "Media".place_id
  WHERE
    "Place".geometry IS NOT NULL
)
SELECT *
FROM 
  places_with_distance
ORDER BY
  distance ASC
LIMIT $3
OFFSET $4;