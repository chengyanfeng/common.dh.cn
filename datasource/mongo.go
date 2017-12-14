package datasource

import (
	"time"

	"github.com/orcaman/concurrent-map"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "common.dh.cn/def"
	. "common.dh.cn/util"
)

var sessionMap = cmap.New()
var EE_MONGO = P{"timeout": 10, "name": "dhe"}
var DH_MONGO = P{"timeout": 10, "name": "dh", "username": "dh", "password": "Pwd0fmE16!@)%"}

type MongoModel struct {
	Cfg    *P
	C      int
	Cname  string
	Query  *P       // find/query condition
	Start  int      // query start at
	Rows   int      // query max rows
	sort   []string // sort
	Select *P       // select field
}

func (m *MongoModel) Session() (_ *mgo.Session, err error) {
	p := *m.Cfg
	key := ToString(p)
	tmp, b := sessionMap.Get(key)
	if !b {
		session, err := mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{ToString(p["host"])},
			Database: ToString(p["name"]),
			Username: ToString(p["username"]),
			Password: ToString(p["password"]),
			Timeout:  time.Duration(ToInt(p["timeout"], 10)) * time.Second,
		})
		if err != nil {
			Error(err)
			return nil, err
		} else {
			sessionMap.Set(key, session)
			return session.Clone(), nil
		}
	} else {
		return tmp.(*mgo.Session).Clone(), nil
	}
}

func (m *MongoModel) Run(collection string, f func(*mgo.Collection)) error {
	session, err := m.Session()
	if err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			Error("Mgo", err)
		}
		session.Close()
	}()
	p := *m.Cfg
	c := session.DB(ToString(p["name"])).C(collection)
	f(c)
	return err
}

func (m *MongoModel) Like(v string) (result interface{}) {
	return &bson.RegEx{Pattern: v, Options: "i"}
}

func (m *MongoModel) Find(p P) *MongoModel {
	m.Query = &p
	return m
}

func (m *MongoModel) Or(ps ...P) *MongoModel {
	q := *m.Query
	tmp := q["$or"]
	or := []P{}
	if tmp != nil {
		or = tmp.([]P)
	}
	for _, p := range ps {
		or = append(or, p)
	}
	q["$or"] = or
	m.Query = &q
	return m
}

func (m *MongoModel) ToString() string {
	return JoinStr(m.Cname, m.Query, m.Start, m.Rows, m.sort, m.Select)
}

func (m *MongoModel) Cache(i ...int) *MongoModel {
	if len(i) > 0 {
		m.C = i[0]
	} else {
		m.C = 3
	}
	return m
}

func (m *MongoModel) Field(s ...string) *MongoModel {
	if m.Select == nil {
		m.Select = &P{}
	}
	for _, k := range s {
		if StartsWith(k, "-") {
			(*m.Select)[k[1:]] = 0
		} else {
			(*m.Select)[k] = 1
		}
	}
	return m
}

func (m *MongoModel) Skip(start int) *MongoModel {
	m.Start = start
	return m
}

func (m *MongoModel) Limit(rows int) *MongoModel {
	m.Rows = rows
	return m
}

func (m *MongoModel) Page(start int, rows int) (total int, list []P) {
	total = m.Count()
	m.Start = start
	m.Rows = rows
	list = m.All()
	return
}

func (m *MongoModel) Sort(s ...string) *MongoModel {
	if !IsEmpty(s) {
		m.sort = s
	}
	return m
}

func (m *MongoModel) All() (r []P) {
	cacheKey := JoinStr(m.ToString(), "All")
	if m.C > 0 {
		tmp := S(cacheKey)
		if tmp != nil {
			r = []P{}
			for _, v := range tmp.([]P) {
				r = append(r, v.Copy())
			}
		}
	}
	if r == nil {
		r = []P{}
		m.Run(m.Cname, func(c *mgo.Collection) {
			q := m.query(c)
			q.All(&r)
		})
		if len(r) > 0 && m.C > 0 {
			S(cacheKey, r, m.C)
		}
	}
	return
}

func (m *MongoModel) One() *P {
	var r *P
	cacheKey := JoinStr(m.ToString(), "One")
	if m.C > 0 {
		tmp := S(cacheKey)
		if tmp != nil {
			r = tmp.(*P)
		}
	}
	if r == nil {
		r = &P{}
		m.Run(m.Cname, func(c *mgo.Collection) {
			q := m.query(c)
			err := q.One(r)
			if err != nil && err.Error() != "not found" {
				Error("One", err)
			}
		})
	}
	if len(*r) > 0 && m.C > 0 {
		S(cacheKey, r, m.C)
	}
	tmp := r.Copy()
	return &tmp
}

func (m *MongoModel) Count() (total int) {
	var tmp interface{}
	cacheKey := JoinStr(m.ToString(), "Count")
	if m.C > 0 {
		tmp = S(cacheKey)
		if tmp != nil {
			Debug("Count from cache", tmp)
			total = tmp.(int)
		}
	}
	if tmp == nil {
		m.Run(m.Cname, func(c *mgo.Collection) {
			q := m.query(c)
			total, _ = q.Count()
		})
	}
	if total > -1 && m.C > 0 {
		S(cacheKey, total, m.C)
	}
	return
}

func (m *MongoModel) Distinct(key string) (r []interface{}) {
	cacheKey := JoinStr(m.ToString(), "Distinct", key)
	if m.C > 0 {
		tmp := S(cacheKey)
		if tmp != nil {
			Debug("Distinct from cache", tmp)
			r = tmp.([]interface{})
		}
	}
	if r == nil {
		r = []interface{}{}
		m.Run(m.Cname, func(c *mgo.Collection) {
			q := m.query(c)
			q.Distinct(key, &r)
		})
	}
	if len(r) > 0 && m.C > 0 {
		S(cacheKey, r, m.C)
	}
	return
}

func (m *MongoModel) Add(docs ...interface{}) (err error) {
	m.Run(m.Cname, func(c *mgo.Collection) {
		if len(docs) == 1 {
			err = c.Insert(docs[0])
		} else {
			err = c.Insert(docs)
		}
	})
	return
}

func (m *MongoModel) Batch(docs []P) (err error) {
	Debug("Batch", JsonEncode(m.Cfg), len(docs))
	tmp := []interface{}{}
	for _, v := range docs {
		v["@ct"] = time.Now()
		tmp = append(tmp, v)
	}
	m.Run(m.Cname, func(c *mgo.Collection) {
		err = c.Insert(tmp...)
		if err != nil {
			Error("Batch", err)
		}
	})
	return
}

func (m *MongoModel) Upsert(selector interface{}, doc interface{}) (err error) {
	m.Run(m.Cname, func(c *mgo.Collection) {
		_, err = c.Upsert(selector, P{"$set": doc})
		if err != nil {
			Error(err)
		}
	})
	return err
}

func (m *MongoModel) Save(p *P) (err error) {
	m.Run(m.Cname, func(c *mgo.Collection) {
		id := (*p)["_id"]
		var oid bson.ObjectId
		switch id.(type) {
		case string:
			oid = bson.ObjectIdHex(id.(string))
		case bson.ObjectId:
			oid = id.(bson.ObjectId)
		}
		(*p)["_id"] = oid
		err = c.UpdateId(oid, P{"$set": p})
		if err != nil {
			Error(err)
		}
	})
	return
}

func (m *MongoModel) Update(selector P, p *P) (err error) {
	m.Run(m.Cname, func(c *mgo.Collection) {
		_, err = c.UpdateAll(selector, P{"$set": p})
		if err != nil {
			Error(err)
		}
	})
	return
}

func (m *MongoModel) Push(selector P, p *P) (err error) {
	m.Run(m.Cname, func(c *mgo.Collection) {
		_, err = c.UpdateAll(selector, P{"$push": p})
		if err != nil {
			Error(err)
		}
	})
	return
}

func (m *MongoModel) RemoveId(id string) {
	m.Run(m.Cname, func(c *mgo.Collection) {
		err := c.RemoveId(bson.ObjectIdHex(id))
		if err != nil && err.Error() != "not found" {
			Error(err)
		}
	})
}

func (m *MongoModel) Remove(selector interface{}) (e error) {
	m.Run(m.Cname, func(c *mgo.Collection) {
		_, err := c.RemoveAll(selector)
		if err != nil {
			Error(err)
			e = err
		}
	})
	return
}

func (m *MongoModel) Drop() (e error) {
	m.Run(m.Cname, func(c *mgo.Collection) {
		err := c.DropCollection()
		if err != nil {
			Error(err)
			e = err
		}
	})
	return
}

func (m *MongoModel) Explain() (result interface{}) {
	p := P{}
	m.Run(m.Cname, func(c *mgo.Collection) {
		q := m.query(c)
		q.Explain(p)
	})
	return p
}

func (m *MongoModel) query(c *mgo.Collection) *mgo.Query {
	q := c.Find(m.Query).Skip(m.Start)
	if m.Rows > 0 {
		q = q.Limit(m.Rows)
	}
	if len(m.sort) > 0 {
		q = q.Sort(m.sort...)
	}
	if m.Select != nil {
		q = q.Select(m.Select)
	}
	return q
}

func D(name string, params ...P) (m *MongoModel) {
	m = &MongoModel{Cname: name}
	if len(params) < 1 {
		p := DH_MONGO
		if EE {
			p = EE_MONGO
		}
		params = []P{p}
	}
	p := params[0]
	m.Cfg = &p
	m.Query = &P{}
	return
}

func (m *MongoModel) Sql(sql string) (r string, err error) {
	r, err = HttpPost(Jdbc_proxy_url, nil, &P{"sql": sql, "db": JsonEncode(m.Cfg)})
	return
}

func (c *MongoModel) Import(tblname string, f func([]P), page ...int) (e error) {
	pos := c.loadPos(tblname)
	data := []P{}
	if IsEmpty(pos) {
		tmp := *D(tblname).Find(P{}).Field("_id").Sort("_id").One()
		pos = ToString(tmp["_id"])
	}
	if !IsEmpty(pos) {
		if len(page) == 0 {
			page = []int{1000}
		}
		p := P{"_id": P{"$gte": ToOid(pos)}}
		Debug("Import", JsonEncode(p))
		tmp := D(tblname).Find(p).Limit(page[0]).Sort("_id").All()
		for _, v := range tmp {
			t := v
			data = append(data, t)
		}
		f(data)
		if len(tmp) > 0 {
			lastRow := tmp[len(data)-1]
			pos = ToString(lastRow["_id"])
			c.savePos(tblname, pos)
		}
	}
	return
}

func (c *MongoModel) loadPos(name string) (r string) {
	p := *D(DbPos).Find(P{"key": c.getStoreKey(name)}).One()
	if len(p) > 0 {
		r = ToString(p["pos"])
	}
	return
}

func (c *MongoModel) savePos(name string, pos string) {
	D(DbPos).Upsert(P{"key": c.getStoreKey(name)}, P{"key": c.getStoreKey(name), "pos": pos})
}

func (c *MongoModel) ClearPos(name string) {
	D(DbPos).Remove(P{"key": c.getStoreKey(name)})
}

func (c *MongoModel) getStoreKey(name string) string {
	return Md5(c.Cfg, name)
}

func MgoLike(v string) (result interface{}) {
	return &bson.RegEx{Pattern: v, Options: "i"}
}

func (m *MongoModel) Pipe(pipe []P) (r []P) {
	r = []P{}
	m.Run(m.Cname, func(c *mgo.Collection) {
		c.Pipe(pipe).All(&r)
	})
	return
}

func (m *MongoModel) Statis(field string) (r P) {
	field = JoinStr("$", field)
	r = P{}
	Debug("Query=", m.Query)
	tmp := m.Pipe([]P{
		{"$match": m.Query},
		{"$group": P{
			"_id": 1,
			"sum": P{"$sum": field},
			"avg": P{"$avg": field},
			"max": P{"$avg": field},
			"min": P{"$avg": field},
		}},
	})
	if len(tmp) > 0 {
		r = tmp[0]
	}
	return
}
