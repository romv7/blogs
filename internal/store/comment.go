package store

func (c *Comment) Save() (err error) {
	cstore := NewCommentStore(c.t)
	err = cstore.Save(c)

	return
}

func (c *Comment) Delete() (err error) {
	cstore := NewCommentStore(c.t)
	err = cstore.Delete(c)

	return
}
