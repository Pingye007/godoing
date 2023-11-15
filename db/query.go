package db

func queryById(id int, instance any) (any, error) {
	out := instance
	_, err := Engine.ID(id).Get(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
