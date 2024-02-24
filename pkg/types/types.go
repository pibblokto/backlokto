package types

type Specs struct {
	Path     string `yaml:"path"`
	RdsArn   string `yaml:"rds_arn"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Target struct {
	Type         string `yaml:"type"`
	S3BucketName string `yaml:"s3_bucket_name"`
	S3BucketKey  string `yaml:"s3_bucket_key"`
	Region       string `yaml:"region"`
	AccessKey    string `yaml:"access_key"`
	SecretKey    string `yaml:"secret_key"`
}

type BackupJob struct {
	Id       string            `yaml:"id"`
	Provider string            `yaml:"provider"`
	Spec     Specs             `yaml:"spec"`
	Targets  []Target          `yaml:"targets"`
	Tags     map[string]string `yaml:"tags"`
}

type Jobs struct {
	Jobs []BackupJob `yaml:"jobs"`
}

type Artifacts struct {
	Filepath string
}
