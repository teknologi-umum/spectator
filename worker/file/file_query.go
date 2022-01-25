package file

func (d *Dependency) IsDebug() bool {
	return d.Environment == "DEVELOPMENT"
}
