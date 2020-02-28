import os, sys, pprint
{% if cookiecutter.run_init != "n" -%}
res = os.popen('go mod vendor && go get -u && go test ./...')
read_res = res.readlines()
pprint.pprint(read_res)
res.close()
sys.exit(0)
print(read_res)
# pprint.pprint(res.readlines())
{% elif cookiecutter.run_init == "n" -%}
{%endif -%}