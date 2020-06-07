#!/bin/sh

mockgen -source domain/item.go -destination domain/mocks/item.go
mockgen -source domain/item_detail.go -destination domain/mocks/item_detail.go
mockgen -source domain/calendar.go -destination domain/mocks/calendar.go
mockgen -source domain/push_token.go -destination domain/mocks/push_token.go
mockgen -source domain/user.go -destination domain/mocks/user.go