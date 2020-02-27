module {{cookiecutter.app_name}}

go 1.13

require (
	github.com/gorilla/mux v1.7.4
	{% if cookiecutter.use_postgres != "n" -%}github.com/lib/pq v1.3.0{%- endif %}
	{% if cookiecutter.use_GORM == "y" -%}github.com/jinzhu/gorm v1.9.11{%- endif %}
	{% if cookiecutter.use_validate_lib == "y" -%}github.com/gookit/validate v1.2.0{%- endif %}
	{% if cookiecutter.use_authboss == "y" -%}github.com/volatiletech/authboss v2.4.0{%- endif %}
	{% if cookiecutter.use_redis == "y" -%}github.com/go-redis/redis/v7 v7.2.0{%- endif %}
	{% if cookiecutter.use_mongo == "y" -%}go.mongodb.org/mongo-driver v1.3.0{%- endif %}
	
)
