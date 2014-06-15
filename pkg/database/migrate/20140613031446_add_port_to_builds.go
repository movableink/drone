package migrate

type rev20140613031446 struct{}

var AddPortToBuilds = &rev20140613031446{}

func (r *rev20140613031446) Revision() int64 {
	return 20140613031446
}

func (r *rev20140613031446) Up(mg *MigrationDriver) error {
	_, err := mg.AddColumn("builds", "port TEXT")
	return err
}

func (r *rev20140613031446) Down(mg *MigrationDriver) error {
	_, err := mg.DropColumns("builds", "port")
	return err
}
