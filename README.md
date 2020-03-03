Template for create go  service

#### Usage:
1) to start create API liccaly use next command
cookiecutter https://github.com/AntonChernov/cookiecutter-go-non-blocking-api.git

2) Avaible choices:

"full_name": "Anton Chernov",
"github_username": "AntonChernov",
"app_name": "mygolangapiservice",
"project_short_description": "Golang API service.",
"use_validate_lib": "y/n",  --> instal lib for validation (github.com/gookit/validate)
"postgres_uri": "y/n", --> here put the DB uri And use or not use the GORM
"use_GORM": "y/n",  --> ORM for go (github.com/jinzhu/gorm)
"use_authboss": "y/n",
"use_redis": "y/n", --> Install lib and not provide the connection
"mongo_uri": "y/n",  --> mongo uri install lib not provide the connection
"run_init": "y/n"  --> Means run commands go mod vendor, go get -u etc. (full list of commants in hooks/postgen.py)


NOTE: Please remember this template not finished and good only for lokal development(But maybe not so gud :) )
If you want improve something Please fill free to fork and fix! 