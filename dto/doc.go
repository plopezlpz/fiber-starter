// Package dto contains the DTOs (request and response objects)
// for this service Rest API or for communicating with other Rest services
//
// In order to avoid writing superfluous mappers it is preferable to use the objects in
// the model with struct tags
// e.g.:
//	package model
// 	type Coffee struct {
//		Id            `json:"id"`
//		Variant	      `json:"variant"`
//		CustomerPrice `json:"price"`
// 		Origin        `json:"origin"`
// 		PurchaedPrice `json:"-"`   // This is only visible in the DB and the business logic, omit from API
//	}
// If it makes sense to write a dto then if request and response are the same, it is ok to write only one dto to
// represent both. It is easy to change case the requirements change.
package dto
