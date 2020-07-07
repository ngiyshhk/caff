package consts

import "github.com/ngiyshhk/caff/model"

const GoModTemplate = `module {{.RepositoryName}}

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.14
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
`

const GoSumTemplate = `cloud.google.com/go v0.26.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
github.com/BurntSushi/toml v0.3.1/go.mod h1:xHWCNGjB5oqiDr8zfno3MHue2Ht5sIBksp03qcyfWMU=
github.com/PuerkitoBio/goquery v1.5.1/go.mod h1:GsLWisAFVj4WgDibEWF4pvYnkVQBpKBKeU+7zCJoLcc=
github.com/andybalholm/cascadia v1.1.0/go.mod h1:GsXiBklL0woXo1j/WYWtSYYC4ouU9PqHO0sqidkEA4Y=
github.com/census-instrumentation/opencensus-proto v0.2.1/go.mod h1:f6KPmirojxKA12rnyqOA5BBL4O983OfeGPqjHWSTneU=
github.com/client9/misspell v0.3.4/go.mod h1:qj6jICC3Q7zFZvVWo7KLAzC3yx5G7kyvSDkc90ppPyw=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/denisenkom/go-mssqldb v0.0.0-20191124224453-732737034ffd/go.mod h1:xbL0rPBG9cCiLr28tMa8zpbdarY27NDyej4t/EjAShU=
github.com/envoyproxy/go-control-plane v0.9.1-0.20191026205805-5f8ba28d4473/go.mod h1:YTl/9mNaCwkRvm6d1a2C3ymFceY/DCBVvsKhRF0iEA4=
github.com/envoyproxy/protoc-gen-validate v0.1.0/go.mod h1:iSmxcyjqTsJpI2R4NaDN7+kN2VEUnK/pcBlmesArF7c=
github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5/go.mod h1:a2zkGnVExMxdzMo3M0Hi/3sEU+cWnZpSni0O6/Yb/P0=
github.com/gin-contrib/sse v0.1.0 h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=
github.com/gin-contrib/sse v0.1.0/go.mod h1:RHrZQHXnP2xjPF+u1gW/2HnVO7nvIa9PG3Gm+fLHvGI=
github.com/gin-gonic/gin v1.6.3 h1:ahKqKTFpO5KTPHxWZjEdPScmYaGtLo8Y4DMHoEsnp14=
github.com/gin-gonic/gin v1.6.3/go.mod h1:75u5sXoLsGZoRN5Sgbi1eraJ4GU3++wFwWzhwvtwp4M=
github.com/go-playground/assert/v2 v2.0.1/go.mod h1:VDjEfimB/XKnb+ZQfWdccd7VUvScMdVu0Titje2rxJ4=
github.com/go-playground/locales v0.13.0 h1:HyWk6mgj5qFqCT5fjGBuRArbVDfE4hi8+e8ceBS/t7Q=
github.com/go-playground/locales v0.13.0/go.mod h1:taPMhCMXrRLJO55olJkUXHZBHCxTMfnGwq/HNwmWNS8=
github.com/go-playground/universal-translator v0.17.0 h1:icxd5fm+REJzpZx7ZfpaD876Lmtgy7VtROAbHHXk8no=
github.com/go-playground/universal-translator v0.17.0/go.mod h1:UkSxE5sNxxRwHyU+Scu5vgOQjsIJAF8j9muTVoKLVtA=
github.com/go-playground/validator/v10 v10.2.0 h1:KgJ0snyC2R9VXYN2rneOtQcw5aHQB1Vv0sFl1UcHBOY=
github.com/go-playground/validator/v10 v10.2.0/go.mod h1:uOYAAleCW8F/7oMFd6aG0GOhaH6EGOAJShg8Id5JGkI=
github.com/go-playground/validator/v10 v10.3.0 h1:nZU+7q+yJoFmwvNgv/LnPUkwPal62+b2xXj0AU1Es7o=
github.com/go-playground/validator/v10 v10.3.0/go.mod h1:uOYAAleCW8F/7oMFd6aG0GOhaH6EGOAJShg8Id5JGkI=
github.com/go-sql-driver/mysql v1.5.0/go.mod h1:DCzpHaOWr8IXmIStZouvnhqoel9Qv2LBy8hT2VhHyBg=
github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe/go.mod h1:8vg3r2VgvsThLBIFL93Qb5yWzgyZWhEmBwUJWevAkK0=
github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b/go.mod h1:SBH7ygxi8pfUlaOkMMuAQtPIUF8ecWP5IEl/CR7VP2Q=
github.com/golang/mock v1.1.1/go.mod h1:oTYuIxOrZwtPieC+H1uAHpcLFnEyAGVDL/k47Jfbm0A=
github.com/golang/protobuf v1.2.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.3.2/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.3.3 h1:gyjaxf+svBWX08ZjK86iN9geUJF0H6gp2IRKX6Nf6/I=
github.com/golang/protobuf v1.3.3/go.mod h1:vzj43D7+SQXF/4pzW/hwtAqwc6iTitCiVSaWz5lYuqw=
github.com/golang/protobuf v1.4.0-rc.1/go.mod h1:ceaxUfeHdC40wWswd/P6IGgMaK3YpKi5j83Wpe3EHw8=
github.com/golang/protobuf v1.4.0-rc.1.0.20200221234624-67d41d38c208/go.mod h1:xKAWHe0F5eneWXFV3EuXVDTCmh+JuBKY0li0aMyXATA=
github.com/golang/protobuf v1.4.0-rc.2/go.mod h1:LlEzMj4AhA7rCAGe4KMBDvJI+AwstrUpVNzEA03Pprs=
github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0/go.mod h1:WU3c8KckQ9AFe+yFwt9sWVRKCVIyN9cPHBJSNnbL67w=
github.com/golang/protobuf v1.4.0/go.mod h1:jodUvKwWbYaEsadDk5Fwe5c77LiNKVO9IDvqG2KuDX0=
github.com/golang/protobuf v1.4.1/go.mod h1:U8fpvMrcmy5pZrNK1lt4xCsGvpyWQ/VVv6QDs8UjoX8=
github.com/golang/protobuf v1.4.2 h1:+Z5KGCizgyZCbGh1KZqA0fcLLkwbsjIzS4aV2v7wJX0=
github.com/golang/protobuf v1.4.2/go.mod h1:oDoupMAO8OvCJWAcko0GGGIgR6R6ocIYbsSw735rRwI=
github.com/google/go-cmp v0.2.0/go.mod h1:oXzfMopK8JAjlY9xF4vHSVASa0yLyX7SntLO5aqRK0M=
github.com/google/go-cmp v0.3.0/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.3.1/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.4.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.5.0/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/gofuzz v1.0.0/go.mod h1:dBl0BpW6vV/+mYPU4Po3pmUjxk6FQPldtuIdl/M65Eg=
github.com/jinzhu/gorm v1.9.14 h1:Kg3ShyTPcM6nzVo148fRrcMO6MNKuqtOUwnzqMgVniM=
github.com/jinzhu/gorm v1.9.14/go.mod h1:G3LB3wezTOWM2ITLzPxEXgSkOXAntiLHS7UdBefADcs=
github.com/jinzhu/inflection v1.0.0 h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=
github.com/jinzhu/inflection v1.0.0/go.mod h1:h+uFLlag+Qp1Va5pdKtLDYj+kHp5pxUVkryuEj+Srlc=
github.com/jinzhu/now v1.0.1/go.mod h1:d3SSVoowX0Lcu0IBviAWJpolVfI5UJVZZ7cO71lE/z8=
github.com/json-iterator/go v1.1.9 h1:9yzud/Ht36ygwatGx56VwCZtlI/2AD15T1X2sjSuGns=
github.com/json-iterator/go v1.1.9/go.mod h1:KdQUCv79m/52Kvf8AW2vK1V8akMuk1QjK/uOdHXbAo4=
github.com/json-iterator/go v1.1.10 h1:Kz6Cvnvv2wGdaG/V8yMvfkmNiXq9Ya2KUv4rouJJr68=
github.com/json-iterator/go v1.1.10/go.mod h1:KdQUCv79m/52Kvf8AW2vK1V8akMuk1QjK/uOdHXbAo4=
github.com/leodido/go-urn v1.2.0 h1:hpXL4XnriNwQ/ABnpepYM/1vCLWNDfUNts8dX3xTG6Y=
github.com/leodido/go-urn v1.2.0/go.mod h1:+8+nEpDfqqsY+g338gtMEUOtuK+4dEMhiQEgxpxOKII=
github.com/lib/pq v1.1.1/go.mod h1:5WUZQaWbwv1U+lTReE5YruASi9Al49XbQIvNi/34Woo=
github.com/mattn/go-isatty v0.0.12 h1:wuysRhFDzyxgEmMf5xjvJ2M9dZoWAXNNr5LSBS7uHXY=
github.com/mattn/go-isatty v0.0.12/go.mod h1:cbi8OIDigv2wuxKPP5vlRcQ1OAZbq2CE4Kysco4FUpU=
github.com/mattn/go-sqlite3 v1.14.0/go.mod h1:JIl7NbARA7phWnGvh0LKTyg7S9BA+6gx71ShQilpsus=
github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 h1:ZqeYNhU3OHLH3mGKHDcjJRFFRrJa6eAM5H+CtDdOsPc=
github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 h1:Esafd1046DLDQ0W1YjYsBW+p8U2u7vzgW2SQVmlNazg=
github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742/go.mod h1:bx2lNnkwVCuqBIxFjflWJWanXIb3RllmbCylyMrvgv0=
github.com/modern-go/reflect2 v1.0.1 h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=
github.com/modern-go/reflect2 v1.0.1/go.mod h1:bx2lNnkwVCuqBIxFjflWJWanXIb3RllmbCylyMrvgv0=
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4/go.mod h1:xMI15A0UPsDsEKsMN9yxemIoYk6Tm2C1GtYGdfGttqA=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
github.com/stretchr/testify v1.4.0/go.mod h1:j7eGeouHqKxXV5pUuKE4zz7dFj8WfuZ+81PSLYec5m4=
github.com/ugorji/go v1.1.7 h1:/68gy2h+1mWMrwZFeD1kQialdSzAb432dtpeJ42ovdo=
github.com/ugorji/go v1.1.7/go.mod h1:kZn38zHttfInRq0xu/PH0az30d+z6vm202qpg1oXVMw=
github.com/ugorji/go/codec v1.1.7 h1:2SvQaVZ1ouYrrKKwoSk2pzd4A9evlKJb9oTL+OaLUSs=
github.com/ugorji/go/codec v1.1.7/go.mod h1:Ax+UKWsSmolVDwsd+7N3ZtXu+yMGCf907BLYF3GoBXY=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
golang.org/x/exp v0.0.0-20190121172915-509febef88a4/go.mod h1:CJ0aWSM057203Lf6IL+f9T1iT9GByDxfZKAQTCR3kQA=
golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3/go.mod h1:UVdnD1Gm6xHRNCYTkRU2/jEulfH38KcIWyp/GAMgvoE=
golang.org/x/lint v0.0.0-20190227174305-5b3e6a55c961/go.mod h1:wehouNa3lNwaWXcvxsM5YxQ5yQlVC4a0KAMCusXpPoU=
golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3/go.mod h1:6SW0HCj/g11FgYtHlgUYUwCkIfeOF89ocIRzGO/8vkc=
golang.org/x/net v0.0.0-20180218175443-cbe0f9307d01/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20180724234803-3673e40ba225/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20180826012351-8a410e7b638d/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20190213061140-3a22650c66bd/go.mod h1:mL1N/T3taQHkDXs73rZJwtUhF3w3ftmwwsq0BUmARs4=
golang.org/x/net v0.0.0-20190311183353-d8887717615a/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20200202094626-16171245cfb2/go.mod h1:z5CRVTTTmAJ677TzLLGU+0bjPO0LkuOLi4/5GtJWs/s=
golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e/go.mod h1:qpuaurCH72eLCgpAm/N6yyVIVM9cpaDIP3A8BGJEC5A=
golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be/go.mod h1:N/0e6XlmueqKjAGxoOufVs8QHGRruUQn6yWY3a++T0U=
golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20181108010431-42b317875d0f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190423024810-112230192c58/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sys v0.0.0-20180830151530-49385e6e1522/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200116001909-b77594299b42/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd h1:xhmwyvizuTgC2qz7ZlMluP20uW+C3Rm0FD/WLDX8884=
golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae h1:Ih9Yo4hSPImZOpfGuA4bR/ORKTAbhZo2AbWNRCnevdo=
golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.2/go.mod h1:bEr9sfX3Q8Zfm5fL9x+3itogRgK3+ptLWKqgva+5dAk=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20190114222345-bf090417da8b/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20190226205152-f727befe758c/go.mod h1:9Yl7xja0Znq3iFh3HoIrodX9oNMXvdceNzlUR8zjMvY=
golang.org/x/tools v0.0.0-20190311212946-11955173bddd/go.mod h1:LCzVGOaR6xXOjkQ3onu1FJEFr0SW1gC7cKk1uF8kGRs=
golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135/go.mod h1:RgjU9mgBXZiqYHBnxXauZ1Gv1EHHAz9KjViQ78xBX0Q=
golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
google.golang.org/appengine v1.1.0/go.mod h1:EbEs0AVv82hx2wNQdGPgUI5lhzA/G0D9YwlJXL52JkM=
google.golang.org/appengine v1.4.0/go.mod h1:xpcJRLb0r/rnEns0DIKYYv+WjYCduHsrkT7/EB5XEv4=
google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8/go.mod h1:JiN7NxoALGmiZfu7CAH4rXhgtRTLTxftemlI0sWmxmc=
google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55/go.mod h1:DMBHOl98Agz4BDEuKkezgsaosCRResVns1a3J2ZsMNc=
google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013/go.mod h1:NbSheEEYHJ7i3ixzK3sjbqSGDJWnxyFXZblF3eUsNvo=
google.golang.org/grpc v1.19.0/go.mod h1:mqu4LbDTu4XGKhr4mRzUsmM4RtVoemTSY81AxZiDr8c=
google.golang.org/grpc v1.23.0/go.mod h1:Y5yQAOtifL1yxbo5wqy6BxZv8vAUGQwXBOALyacEbxg=
google.golang.org/grpc v1.27.0/go.mod h1:qbnxyOmOxrQa7FizSgH+ReBfzJrCY1pSN7KXBS8abTk=
google.golang.org/protobuf v0.0.0-20200109180630-ec00e32a8dfd/go.mod h1:DFci5gLYBciE7Vtevhsrf46CRTquxDuWsQurQQe4oz8=
google.golang.org/protobuf v0.0.0-20200221191635-4d8936d0db64/go.mod h1:kwYJMbMJ01Woi6D6+Kah6886xMZcty6N08ah7+eCXa0=
google.golang.org/protobuf v0.0.0-20200228230310-ab0ca4ff8a60/go.mod h1:cfTl7dwQJ+fmap5saPgwCLgHXTUD7jkjRqWcaiX5VyM=
google.golang.org/protobuf v1.20.1-0.20200309200217-e05f789c0967/go.mod h1:A+miEFZTKqfCUM6K7xSMQL9OKL/b6hQv+e19PK+JZNE=
google.golang.org/protobuf v1.21.0/go.mod h1:47Nbq4nVaFHyn7ilMalzfO3qCViNmqZ2kzikPIcrTAo=
google.golang.org/protobuf v1.22.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.23.0 h1:4MY060fB1DLGMB/7MBTLnwQUY6+F09GEiz6SsrNqyzM=
google.golang.org/protobuf v1.23.0/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.23.1-0.20200526195155-81db48ad09cc/go.mod h1:EGpADcykh3NcUnDUJcl1+ZksZNG86OlYog2l/sGQquU=
google.golang.org/protobuf v1.25.0 h1:Ejskq+SyPohKW+1uil0JJMtmHCgJPJ/qWTxr8qp+R4c=
google.golang.org/protobuf v1.25.0/go.mod h1:9JNX74DMeImyA3h4bdi1ymwjUzf21/xIlbajtzgsN7c=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.2.8 h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=
gopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.3.0 h1:clyUAQHOM3G0M3f5vQj7LuJrETvjVot3Z5el9nffUtU=
gopkg.in/yaml.v2 v2.3.0/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
honnef.co/go/tools v0.0.0-20190102054323-c2f93a96b099/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc/go.mod h1:rf3lG4BRIbNafJWhAfAdb/ePZxsR/4RtNHQocxwk9r4=
`

const ConstsContextKeyTemplate = `package consts

// ContextKey .
type ContextKey = int

// ContextDB .
const ContextDB ContextKey = iota
`

const ModelMysqlTemplate = `package model

import "fmt"

type Mysql struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DBName string
}

func NewMysql(user, pass, host string, port int, dbName string) *Mysql {
	return &Mysql{
		User:   user,
		Pass:   pass,
		Host:   host,
		Port:   port,
		DBName: dbName,
	}
}

func (m *Mysql) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.User, m.Pass, m.Host, m.Port, m.DBName)
}
`

const UtilContextTemplate = `package utils

import (
	"context"
	"{{.RepositoryName}}/consts"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// ContextWithDB .
func ContextWithDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, consts.ContextDB, db)
}

// GetDBFromContext .
func GetDBFromContext(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value(consts.ContextDB).(*gorm.DB)
	if !ok {
		return nil, errors.New("DB is not found in context")
	}
	if err := db.Error; err != nil {
		return nil, err
	}
	return db, nil
}
`

const InfraMysqlTemplate = `package infrastructure

import "github.com/jinzhu/gorm"

var db *gorm.DB

func InitDB(config string) error {
	var err error
	db, err = gorm.Open("mysql", config)
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
`

const IGatewayTemplate = `package {{.LastIGatewayPackageName}}

import (
	"context"
	"{{.RepositoryName}}/{{.ModelPackageName}}"
)

type {{.SchemaName}} interface {
	Get(ctx context.Context, id string) (*{{.ModelPackageName}}.{{.SchemaName}}, error)
	List(ctx context.Context) ([]*{{.ModelPackageName}}.{{.SchemaName}}, error)
	Create(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error
	Update(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error
	Delete(ctx context.Context, id string) error
}
`

const GatewayTemplate = `package {{.LastGatewayPackageName}}

import (
	"context"
	"{{.RepositoryName}}/{{.IGatewayPackageName}}"
	"{{.RepositoryName}}/{{.ModelPackageName}}"
	"{{.RepositoryName}}/utils"
)

// {{.SchemaName}} .
type {{.SchemaName}} struct {
}

// New{{.SchemaName}} .
func New{{.SchemaName}}() {{.IGatewayPackageName}}.{{.SchemaName}} {
	return &{{.SchemaName}}{}
}

// Get .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Get(ctx context.Context, id string) (*{{.ModelPackageName}}.{{.SchemaName}}, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	result := &{{.ModelPackageName}}.{{.SchemaName}}{}
	err = db.Where("id = ?", id).First(result).Error
	return result, err
}

// List .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) List(ctx context.Context) ([]*{{.ModelPackageName}}.{{.SchemaName}}, error) {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var result []*{{.ModelPackageName}}.{{.SchemaName}}
	err = db.Find(&result).Error
	return result, err
}

// Create .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Create(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Create({{.LowerSchemaName}}).Error
}

// Update .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Update(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Update({{.LowerSchemaName}}).Error
}

// Delete .
func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Delete(ctx context.Context, id string) error {
	db, err := utils.GetDBFromContext(ctx)
	if err != nil {
		return err
	}
	return db.Where("id = ?", id).Delete(&{{.ModelPackageName}}.{{.SchemaName}}{}).Error
}
`

const IControllerTemplate = `package {{.LastIControllerPackageName}}

import "context"

type {{.SchemaName}} interface {
	Get(c context.Context)
	List(c context.Context)
	Post(c context.Context)
	Put(c context.Context)
	Delete(c context.Context)
}
`

const ControllerTemplate = `package {{.LastControllerPackageName}}

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"{{.RepositoryName}}/{{.IControllerPackageName}}"
	"{{.RepositoryName}}/infrastructure"
	"{{.RepositoryName}}/{{.IUsecasePackageName}}"
	"{{.RepositoryName}}/utils"
)

type {{.SchemaName}} struct {
	{{.LowerSchemaName}}Usecase {{.IUsecasePackageName}}.{{.SchemaName}}
}

func New{{.SchemaName}}({{.LowerSchemaName}}Usecase {{.IUsecasePackageName}}.{{.SchemaName}}) {{.IControllerPackageName}}.{{.SchemaName}} {
	return &{{.SchemaName}}{ {{.LowerSchemaName}}Usecase: {{.LowerSchemaName}}Usecase }
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Get(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	ctx := utils.ContextWithDB(gc, infrastructure.GetDB())
	{{.LowerSchemaName}}, err := {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Usecase.Get(ctx, gc.Param("id"))
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, {{.LowerSchemaName}})
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) List(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	ctx := utils.ContextWithDB(gc, infrastructure.GetDB())
	{{.LowerSchemaName}}s, err := {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Usecase.List(ctx)
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, {{.LowerSchemaName}}s)
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Post(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Usecase.Create(ctx, nil)
	})
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(201, gin.H{
		"message": "created",
	})
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Put(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Usecase.Update(ctx, nil)
	})
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, gin.H{
		"message": "updated",
	})
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Delete(ic context.Context) {
	gc, ok := ic.(*gin.Context)
	if !ok {
		gc.JSON(500, gin.H{
			"message": "failed to get *gin.Context",
		})
		return
	}

	err := infrastructure.GetDB().Transaction(func(tx *gorm.DB) error {
		ctx := utils.ContextWithDB(gc, tx)
		return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Usecase.Delete(ctx, gc.Param("id"))
	})
	if err != nil {
		gc.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gc.JSON(200, gin.H{
		"message": "deleted",
	})
}
`

const IUsecaseTemplate = `package {{.LastIUsecasePackageName}}

import (
	"context"
	"{{.RepositoryName}}/{{.ModelPackageName}}"
)

type {{.SchemaName}} interface {
	Get(ctx context.Context, id string) (*{{.ModelPackageName}}.{{.SchemaName}}, error)
	List(ctx context.Context) ([]*{{.ModelPackageName}}.{{.SchemaName}}, error)
	Create(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error
	Update(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error
	Delete(ctx context.Context, id string) error
}
`

const UsecaseTemplate = `package {{.LastUsecasePackageName}}

import (
	"context"
	"{{.RepositoryName}}/{{.IGatewayPackageName}}"
	"{{.RepositoryName}}/{{.IUsecasePackageName}}"
	"{{.RepositoryName}}/{{.ModelPackageName}}"
)

type {{.SchemaName}} struct {
	{{.LowerSchemaName}}Gateway {{.IGatewayPackageName}}.{{.SchemaName}}
}

func New{{.SchemaName}}({{.LowerSchemaName}}Gateway {{.IGatewayPackageName}}.{{.SchemaName}}) {{.IUsecasePackageName}}.{{.SchemaName}} {
	return &{{.SchemaName}}{ {{.LowerSchemaName}}Gateway: {{.LowerSchemaName}}Gateway }
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Get(ctx context.Context, id string) (*{{.ModelPackageName}}.{{.SchemaName}}, error) {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.Get(ctx, id)
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) List(ctx context.Context) ([]*{{.ModelPackageName}}.{{.SchemaName}}, error) {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.List(ctx)
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Create(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.Create(ctx, {{.LowerSchemaName}})
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Update(ctx context.Context, {{.LowerSchemaName}} *{{.ModelPackageName}}.{{.SchemaName}}) error {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.Update(ctx, {{.LowerSchemaName}})
}

func ({{.LowerInitialSchemaName}} *{{.SchemaName}}) Delete(ctx context.Context, id string) error {
	return {{.LowerInitialSchemaName}}.{{.LowerSchemaName}}Gateway.Delete(ctx, id)
}
`

var Templates = []*model.Template{
	{
		Type:  ControllerTemplateType,
		Value: ControllerTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.ControllerPackageName
		},
	},
	{
		Type:  IControllerTemplateType,
		Value: IControllerTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.IControllerPackageName
		},
	},
	{
		Type:  GatewayTemplateType,
		Value: GatewayTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.GatewayPackageName
		},
	},
	{
		Type:  IGatewayTemplateType,
		Value: IGatewayTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.IGatewayPackageName
		},
	},
	{
		Type:  UsecaseTemplateType,
		Value: UsecaseTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.UsecasePackageName
		},
	},
	{
		Type:  IUsecaseTemplateType,
		Value: IUsecaseTemplate,
		PackageNameFunc: func(schema *model.Schema) string {
			return schema.IUsecasePackageName
		},
	},
}
