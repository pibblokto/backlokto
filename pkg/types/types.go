package types

type Specs struct {
	Path             string `yaml:"path"`
	RdsArn           string `yaml:"rds_arn"`
	ConnectionString string `yaml:"connection_string"`
	//Extension                string `yaml:"extension"`
	// rdsArn string
	// connectionString string
	// ...
}

type TargetSpecs struct {
	S3BucketName string `yaml:"s3_bucket_name"`
	S3BucketKey  string `yaml:"s3_bucket_key"`
	// Local string `yaml:"local"`
	// ...
}

type Target struct {
	Type       string      `yaml:"type"`
	TargetSpec TargetSpecs `yaml:"target_spec"`
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

//var TargetsMap = map[string]func(confparse.Target){
//	// "s3": target.S3Target
//	// "local": target.LocalTarget
//}
