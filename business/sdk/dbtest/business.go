package dbtest

import (
	"time"

	"github.com/Klimentin0/courses-service1/business/domain/homebus"
	"github.com/Klimentin0/courses-service1/business/domain/homebus/stores/homedb"
	"github.com/Klimentin0/courses-service1/business/domain/productbus"
	"github.com/Klimentin0/courses-service1/business/domain/productbus/stores/productdb"
	"github.com/Klimentin0/courses-service1/business/domain/userbus"
	"github.com/Klimentin0/courses-service1/business/domain/userbus/stores/usercache"
	"github.com/Klimentin0/courses-service1/business/domain/userbus/stores/userdb"
	"github.com/Klimentin0/courses-service1/business/domain/vproductbus"
	"github.com/Klimentin0/courses-service1/business/domain/vproductbus/stores/vproductdb"
	"github.com/Klimentin0/courses-service1/business/sdk/delegate"
	"github.com/Klimentin0/courses-service1/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// BusDomain represents all the business domain apis needed for testing.
type BusDomain struct {
	Delegate *delegate.Delegate
	Home     *homebus.Business
	Product  *productbus.Business
	User     *userbus.Business
	VProduct *vproductbus.Business
}

func newBusDomains(log *logger.Logger, db *sqlx.DB) BusDomain {
	delegate := delegate.New(log)
	userBus := userbus.NewBusiness(log, delegate, usercache.NewStore(log, userdb.NewStore(log, db), time.Hour))
	productBus := productbus.NewBusiness(log, userBus, delegate, productdb.NewStore(log, db))
	homeBus := homebus.NewBusiness(log, userBus, delegate, homedb.NewStore(log, db))
	vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(log, db))

	return BusDomain{
		Delegate: delegate,
		Home:     homeBus,
		Product:  productBus,
		User:     userBus,
		VProduct: vproductBus,
	}
}
