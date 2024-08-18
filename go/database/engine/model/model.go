package model

type Data struct {
	Engine          string
	ConnStr         string
	Queries         string
	BorderStyle     int
	DestFolder      string
	LuaTabulaPath   string
	TabulaLogFile   string
	HeaderStyleLink string
	Option          Option
}

type Option int

const (
	RUN Option = iota + 1
	TABLES
    TABLE_INFO
)

func (Data) GetTableInfoQuery(tableName string) string {
	return `SELECT 
                UPPER(CAST(c.column_name AS VARCHAR)) AS column_name,
                c.data_type,
                CASE
                    WHEN c.is_nullable = 'YES' THEN ' '
                    ELSE ' '
                END AS not_null,
                CASE
                    WHEN c.character_maximum_length IS NULL THEN '-'
                    ELSE CAST(c.character_maximum_length AS VARCHAR)
                END AS length,
                CASE  
                    WHEN tc.constraint_type = 'PRIMARY KEY' THEN '  PRIMARY KEY'
                    WHEN tc.constraint_type = 'FOREIGN KEY' THEN '  FOREIGN KEY'
                    ELSE '-'
                END AS constraint_type,
                CASE 
                    WHEN tc.constraint_type = 'FOREIGN KEY' THEN 
                       '  ' || kcu2.table_name || '.' || kcu2.column_name
                    ELSE 
                        '-'
                END AS referenced_table_column
                FROM 
                    information_schema.columns AS c
                LEFT JOIN 
                    information_schema.key_column_usage AS kcu 
                    ON c.column_name = kcu.column_name 
                    AND c.table_name = kcu.table_name
                LEFT JOIN 
                    information_schema.table_constraints AS tc 
                    ON kcu.constraint_name = tc.constraint_name 
                    AND kcu.table_name = tc.table_name
                LEFT JOIN 
                    information_schema.referential_constraints AS rc 
                    ON tc.constraint_name = rc.constraint_name 
                    AND tc.table_schema = rc.constraint_schema
                LEFT JOIN 
                    information_schema.key_column_usage AS kcu2 
                    ON rc.unique_constraint_name = kcu2.constraint_name 
                    AND rc.unique_constraint_schema = kcu2.table_schema
                WHERE 
                    c.table_name = '`+ tableName + `';`
}
