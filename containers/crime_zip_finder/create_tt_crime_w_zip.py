#!/usr/bin/env python
# coding: utf-8

import psycopg2
import pandas as pd
from shapely.geometry import shape, Point
import json
import os

host = "10.100.138.27"  # postgresql on AWS host
port = "5432"
user = os.getenv('DB_USER')
password = os.getenv('DB_PASSWORD')
dbname = "fullstack_api"


from sqlalchemy import create_engine
engine = create_engine("postgresql://{0}:{1}@{2}:{3}/{4}".format(user, password, host, port, dbname))

db_connection = psycopg2.connect(host=host,
                                 dbname=dbname, 
                                 user=user,
                                 password=password
 )

cursor = db_connection.cursor()

s = 'CREATE TABLE IF NOT EXISTS t_crime_data_wzip ( '
s += 'id serial PRIMARY KEY, '
s += 'case_number VARCHAR, '
s += 'date DATE, '
s += 'block VARCHAR, '
s += 'iucr VARCHAR, '
s += 'primary_type VARCHAR, '
s += 'description VARCHAR, '
s += 'location_description VARCHAR, '
s += 'arrest BOOLEAN, '
s += 'domestic BOOLEAN, '
s += 'beat VARCHAR, '
s += 'district VARCHAR, '
s += 'ward NUMERIC, '
s += 'community_area VARCHAR, '
s += 'fbi_code VARCHAR, '
s += 'x_coordinate NUMERIC, '
s += 'y_coordinate NUMERIC, '
s += 'year NUMERIC, '
s += 'updated_on DATE, '
s += 'latitude NUMERIC, '
s += 'longitude NUMERIC, '
s += 'zip VARCHAR);'

cursor.execute(s)
db_connection.commit()

s = 'DELETE FROM t_crime_data_wzip '
s += 'WHERE id = ANY(ARRAY( '
s += 'SELECT id FROM t_crime_data_wzip '
s += 'ORDER BY id DESC LIMIT 500000));'

try:
    cursor.execute(s)
    db_connection.commit()
except:
    pass

def createList(r1, r2):
    return [item for item in range(r1, r2+1)]

s = "SELECT id " 
s += "FROM t_crime_data_wzip "
s += "ORDER BY id DESC "
s += "LIMIT 1;"

cursor.execute(s)

offset = cursor.fetchone()

offset = offset[0]

s = "SELECT id " 
s += "FROM t_crime_data "
s += "ORDER BY id DESC "
s += "LIMIT 1;"

cursor.execute(s)

max_crime_id = cursor.fetchone()


s = "SELECT id, case_number, date, block, iucr, primary_type, description, "
s += "location_description, arrest, domestic, beat, district, ward, community_area, "
s += "fbi_code, x_coordinate, y_coordinate, year, updated_on, latitude::float, longitude::float "
s += "FROM t_crime_data "
s += "WHERE id > {0} AND id < {1};".format(offset, max_crime_id) # change force id after manual load

cursor.execute(s)

rows=cursor.fetchall()

with open('zip_boundaries_with_ca.geojson') as fp:
        boundaries = json.load(fp)

boundaries['features'][0]['properties']['area_numbe']

def find_zip(rows):
    new_rows = []
    for row in rows:
        temp_tupl = row
        temp_list = list(temp_tupl)
        temp_list.append(None)
        do_f = 0
        for feature in boundaries['features']:
            com_area = feature['properties']['area_numbe']
            zip_code = feature['properties']['zip']
            if row[13] == com_area:
                polygon = shape(feature["geometry"])
                try:
                    point1 = Point(row[20], row[19])
                except:
                    pass
                try:
                    if polygon.contains(point1):
                        temp_list[21] = zip_code
                        do_f = 1
                    else:
                        pass
                except:
                    pass
        if do_f == 0:
            for feature in boundaries['features']:
                com_area = feature['properties']['area_numbe']
                zip_code = feature['properties']['zip']
                polygon = shape(feature["geometry"])
                try:
                    point1 = Point(row[20], row[19])
                except:
                    continue
                try:
                    if polygon.contains(point1):
                        temp_list[21] = zip_code
                        do_f = 1
                    else:
                        continue
                except:
                    continue
        new_tuple = tuple(temp_list)
        new_rows.append(new_tuple)
    return new_rows


rows_with_zip = find_zip(rows)

crimedata_df = pd.DataFrame(rows_with_zip, 
                            columns = [
                                'id', 
                                'case_number', 
                                'date', 
                                'block',
                                'iucr', 
                                'primary_type', 
                                'description',
                                'location_description',
                                'arrest',
                                'domestic',
                                'beat',
                                'district',
                                'ward',
                                'community_area',
                                'fbi_code',
                                'x_coordinate',
                                'y_coordinate',
                                'year',
                                'updated_on',
                                'latitude',
                                'longitude',
                                'zip'
                            ]
                            )

crimedata_df.to_sql('t_crime_data_wzip', con=engine, if_exists='append', index=False)

cursor.close()
db_connection.close()