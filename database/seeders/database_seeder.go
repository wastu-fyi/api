package seeders

type DatabaseSeeder struct {
}

func (s *DatabaseSeeder) Signature() string {
	return "DatabaseSeeder"
}

func (s *DatabaseSeeder) Run() error {
	return nil
}
