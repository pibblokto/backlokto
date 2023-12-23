package types

type Specs struct {
	Path     string `yaml:"path"`
	RdsArn   string `yaml:"rds_arn"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Target struct {
	Type         string `yaml:"type"`
	S3BucketName string `yaml:"s3_bucket_name"`
	S3BucketKey  string `yaml:"s3_bucket_key"`
}

type BackupJob struct {
	Id       string   `yaml:"id"`
	Provider string   `yaml:"provider"`
	Spec     Specs    `yaml:"spec"`
	Targets  []Target `yaml:"targets"`
}

type Jobs struct {
	Jobs []BackupJob `yaml:"jobs"`
}
