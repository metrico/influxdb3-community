from flightsql import connect, FlightSQLClient

client = FlightSQLClient(host='localhost',port=8082,insecure=True,metadata={'bucket':'company_sensors'})
conn = connect(client)
cursor = conn.cursor()
cursor.execute('SELECT * FROM cpu LIMIT 1;')
# print("columns:", cursor.description)
print("rows:", [r for r in cursor])
