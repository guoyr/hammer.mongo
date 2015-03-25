package profiles

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	basic insert only profile, insert a doc of 400 bytes
*/

type customerWorkloadProfile struct {
	UID int64

	indexGroup bool

	initProfile sync.Once

	session *mgo.Session
}

var _customerWorkloadProfile customerWorkloadProfile

// func Int2ObjId(i int64) string {
// 	// return string represenation of UID
// }

func (i customerWorkloadProfile) SendNext(s *mgo.Session, worker_id int) error {
	var err error
	c := s.DB(getDBName("test")).C(getCollectionName("users"))

	zip := rand.Intn(90000) + 10000

	var result = make(map[string]interface{})

	err = c.Find(bson.M{"zip": zip}).One(result)
	// err = c.Database.Run(bson.D{{"find", c.Name}, {"zip",zip}}, result)
	log.Println(result)

	// panicOnError(err)
	return err
}

func initCustomerTest(session *mgo.Session, _initdb bool) {
	log.Println("Initialize simple DB. with initdb flag ", _initdb)

	_initdb = true

	if !_initdb {
		panic("flag is false")
	}

	// drop the colelction here  TODO?  FIXME:
	if _initdb {
		log.Println(". Init DB, drop collections")
		session.DB(_db_name).C("people").DropCollection()
		// may drop DB here as well TODO:
	} // this will be moved to each profile. FIXME:

	for i := 1; i < _multi_db; i++ {
		for j := 1; j < _multi_col; j++ {
			collection := session.DB(default_db_name_prefix + strconv.Itoa(i)).C(default_col_name_prefix + strconv.Itoa(j))
			err := collection.EnsureIndexKey("name")
			if err != nil {
				panic(err)
			}

			// err = collection.EnsureIndexKey("group")
			err = collection.EnsureIndexKey("uid")
			if err != nil {
				panic(err)
			}
		}
	}
}

func (i customerWorkloadProfile) SetupTest(s *mgo.Session, _initdb bool) error {
	i.session = s

	f := func() {
		initCustomerTest(s, _initdb)
	}

	_customerWorkloadProfile.initProfile.Do(f)
	return nil
}

func (i customerWorkloadProfile) CsvString(total_time float64) string {
	return ""
}

func (i customerWorkloadProfile) CsvHeader() string {
	return ""
}

func init() {
	// fmt.Println("Init INSERT profile")

	atomic.StoreInt64(&_customerWorkloadProfile.UID, -1) // UID starts with 1

	registerProfile("CUSTOMER", func() Profile {
		return Profile(_customerWorkloadProfile) // use the same instance
	})

	s := os.Getenv("HT_INDEX_GROUP")
	if s == "" {
		_customerWorkloadProfile.indexGroup = false
	} else {
		_customerWorkloadProfile.indexGroup = true
	}

	// fmt.Println("Done Init INSERT profile")
}
