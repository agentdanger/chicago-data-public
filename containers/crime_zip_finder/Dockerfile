FROM python:3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN pip install psycopg2 
RUN pip install pandas 
RUN pip install shapely 
RUN pip install SQLAlchemy
CMD ["/app/main"]