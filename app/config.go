package main

import (
	"io/ioutil"
)

import (
	"gopkg.in/yaml.v2"
)

// ConfYaml is config structure.
type ConfYaml struct {
	Core  SectionCore  `yaml:"core"`
	Kafka SectionKafka `yaml:"kafka"`
	Es    SectionEs    `yaml:"es"`
}

// SectionPID is sub section of config.
type SectionPID struct {
	Enabled  bool   `yaml:"enabled"`
	Path     string `yaml:"path"`
	Override bool   `yaml:"override"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	FailFastTimeout int        `yaml:"fail_fast_timeout"`
	WorkerNum       int64      `yaml:"worker_num"`
	QueueNum        int64      `yaml:"queue_num"`
	PID             SectionPID `yaml:"pid"`
}

// SectionKafka is sub section of config.
type SectionKafka struct {
	Brokers       string `yaml:"brokers"`
	Topic         string `yaml:"topic"`
	ConsumerGroup string `yaml:"consumer_group"`
}

type SectionEs struct {
	EsHosts         []string `yaml:"es_hosts"`
	ShardNum        int32    `yaml:"shard_num"`
	ReplicaNum      int32    `yaml:"replica_num"`
	RefreshInterval int32    `yaml:"refresh_interval"`

	Index                 string `yaml:"index"`
	IndexTimeSuffixFormat string `yaml:"index_time_suffix_format"`
	Type                  string `yaml:"type"`
	KibanaTimeField       string `yaml:"kibana_time_filed"`
	KibanaTimeFormat      string `yaml:"kibana_time_format"`
	BulkSize              int32  `yaml:"bulk_size"`
	BulkTimeout           int32  `yaml:"bulk_timeout"`
}

// LoadConfYaml provide load yml config.
func LoadConfYaml(confPath string) (ConfYaml, error) {
	var config ConfYaml

	configFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return config, err
	}

	return config, nil
}
