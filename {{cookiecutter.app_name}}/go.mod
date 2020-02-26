module {{cookiecutter.app_name}}

go 1.13

require (
	github.com/gorilla/mux v1.7.4
	{% if cookiecutter.use_postgres != "n" -%}github.com/lib/pq v1.3.0{%- endif %}
	{% if cookiecutter.use_GORM == "y" -%}github.com/jinzhu/gorm v1.9.11{%- endif %}
)
