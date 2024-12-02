package sql

func SearchPlacesBaseQuery() string {
	return `
  SELECT
    id,
    name,
    state,
    city,
    zip_code AS code,
    address_line1 AS address1,
    address_line2 AS address2,
    CEIL(ST_Distance(
      ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
      geometry::geography
    ) / 1000) AS distance -- メートル単位からキロメートル単位に変換
  FROM 
    "Place"
  WHERE
  (
    -- 検索キーワード条件
    %s
  )
  AND
  geometry IS NOT NULL
  ORDER BY distance
  LIMIT 10;
  `
}
