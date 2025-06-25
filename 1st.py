# Databricks notebook source
# MAGIC %md
# MAGIC ### pyspark tutorial
# MAGIC

# COMMAND ----------

from pyspark.sql import SparkSession
spark = SparkSession.builder.appName("Spark DataFrames").getOrCreate()
data=[(1,'praveen'),(2,'yogesh'),(3,'suresh'),(4,'ramesh'),(5,'suresh'),(6,'ramesh'),(7,'ramesh'),(8,'ramesh'),(9,'ramesh'),(10,'ramesh')]

schema=['id','name']
df=spark.createDataFrame(data=data,schema=schema)
df.show(3,truncate=False)

# COMMAND ----------

df=spark.read.option('header','True').option('inferSchema','True').csv('/Volumes/workspace/default/praveen/bigmart.csv')

# COMMAND ----------

display(df)

# COMMAND ----------

from pyspark.sql import SparkSession
from pyspark.sql.functions import col,lit,concat_ws
spark = SparkSession.builder.appName("Spark DataFrames").getOrCreate()
data =([(1,'praveen'),(2,'yogesh'),(3,'suresh'),(4,'weds'),(5,'suresh'),(6,'ramesh'),(7,'ramesh'),(8,'ramesh'),(9,'ramesh')])
schema=['id','name']
df=spark.createDataFrame(data=data,schema=schema)

df2=df.withColumn("nigesh",lit(1000)+col("id"))
df3=df2.withColumn("kani",concat_ws(" ",col("name"),col('id')))

df5=[col for col in df.columns]
df.select(df.columns[0:1]).display()


# COMMAND ----------

from pyspark.sql import SparkSession
from pyspark.sql.types import *
spark=SparkSession.builder.appName("Spark DataFrames").getOrCreate()
data=[(1,[{'f':"praveen",'l':"mumthi"}])]
schema=StructType([
  StructField("nos",IntegerType(),True),
  StructField("array",ArrayType(StructType([
    StructField("f",StringType(),True),
    StructField("l",StringType(),True)
  ])),True)
])
df=spark.createDataFrame(data,schema)
display(df)