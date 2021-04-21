package datastore

import ()

// DatastoreInterface methods are implemented by any concrete datastore
type DatastoreInterface interface {
	Connect() bool
	Disconnect()
}
