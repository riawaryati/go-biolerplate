#mock infra
mockgen -package=mock_infra -source=infra/db.go -destination=infra/mock/db.go

#mock repo
mockgen -package=master -source=repo/master/city.go -destination=mock/repo/master/mock_city_repo.go
mockgen -package=master -source=repo/master/country.go -destination=mock/repo/master/mock_country_repo.go