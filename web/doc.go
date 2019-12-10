// Package web provides web layer API over service layer.
//
// Why put web stuff in a separate package?
//
// The idea is to minimize loading a lot of logic into handlers.
// The only thing handlers have to know about service logic is the specific service
// interface handling the request. Handlers only responsibility is to translate
// http request to domain/service request. They should serve as a thin layer.
package web
